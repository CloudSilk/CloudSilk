// Package middleware 提供 CloudSilk 业务层的接口级权限守卫。
//
// 背景：AuthRequired 已完成认证与 URL 级鉴权（usercenter Casbin），
// 但敏感业务操作（作废排程、完成检验、出库等）需要额外的角色门槛。
// 守卫策略：超级管理员角色（config superAdminRoleID）始终放行；
// 其余角色需命中系统参数中配置的角色白名单；参数未配置时默认拒绝（安全缺省）。
package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/constants"
	usercenter "github.com/CloudSilk/usercenter/proto"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// 权限参数在系统参数表中的模块代号
const permissionParamsCode = "permission"

// paramCacheTTL 参数缓存时长：角色白名单变更后最迟此时间内生效
const paramCacheTTL = time.Minute

var (
	paramCache sync.Map // key -> cachedParam
)

type cachedParam struct {
	value   string
	expires time.Time
}

// getParamValue 读取系统参数（带短缓存，避免每个请求查库）
func getParamValue(key string) string {
	if c, ok := paramCache.Load(key); ok {
		cp := c.(cachedParam)
		if time.Now().Before(cp.expires) {
			return cp.value
		}
	}
	m := &model.SystemParamsConfig{}
	err := model.DB.DB().Where("`code` = ? AND `key` = ?", permissionParamsCode, key).First(m).Error
	value := ""
	if err == nil {
		value = m.Value
	}
	paramCache.Store(key, cachedParam{value: value, expires: time.Now().Add(paramCacheTTL)})
	return value
}

// hasAnyRole 判断用户角色与允许集合是否有交集
func hasAnyRole(userRoles, allowed []string) bool {
	if len(userRoles) == 0 || len(allowed) == 0 {
		return false
	}
	set := make(map[string]struct{}, len(allowed))
	for _, r := range allowed {
		set[strings.TrimSpace(r)] = struct{}{}
	}
	for _, r := range userRoles {
		if _, ok := set[strings.TrimSpace(r)]; ok {
			return true
		}
	}
	return false
}

// RequireRoles 返回只允许指定角色（或超管）访问的 gin 中间件。
// allowedRoles 为空时仅超管可访问。
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, user := ucmiddleware.GetUser(c)
		if !ok || user == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": 403, "message": "未获取到用户信息",
			})
			return
		}
		roles := userRoles(user)
		if constants.SuperAdminRoleID != "" && hasAnyRole(roles, []string{constants.SuperAdminRoleID}) {
			c.Next()
			return
		}
		if hasAnyRole(roles, allowedRoles) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code": 403, "message": "当前角色无权执行此操作",
		})
	}
}

// RequirePermissionParam 与 RequireRoles 类似，但允许角色来自系统参数
// （code=permission, key=paramKey），值形如 "roleId1,roleId2"；参数为空时仅超管可访问。
func RequirePermissionParam(paramKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, user := ucmiddleware.GetUser(c)
		if !ok || user == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": 403, "message": "未获取到用户信息",
			})
			return
		}
		roles := userRoles(user)
		if constants.SuperAdminRoleID != "" && hasAnyRole(roles, []string{constants.SuperAdminRoleID}) {
			c.Next()
			return
		}
		value := getParamValue(paramKey)
		if value == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": 403, "message": "此操作未配置授权角色，请联系管理员（系统参数 permission/" + paramKey + "）",
			})
			return
		}
		allowed := strings.Split(value, ",")
		if hasAnyRole(roles, allowed) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code": 403, "message": "当前角色无权执行此操作",
		})
	}
}

func userRoles(user *usercenter.CurrentUser) []string {
	if user == nil {
		return nil
	}
	return user.RoleIDs
}
