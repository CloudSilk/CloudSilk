import umiRequest from 'umi-request';
import { getToken, GetDetailResponse, User } from '@swiftease/atali-pkg';

/**
 * 认证相关 API（统一 service 层，供 app.tsx 与 AvatarDropdown 等调用）
 */

const authHeaders = (): Record<string, string> => ({
  'Content-Type': 'application/json',
  authorization: 'Bearer ' + getToken(),
});

/** 获取当前登录用户信息 */
export async function queryCurrentUser(options?: { [key: string]: any }) {
  return umiRequest<GetDetailResponse<User.UserProfile>>('/api/core/auth/user/profile', {
    method: 'GET',
    headers: authHeaders(),
    ...(options || {}),
  });
}

/** 退出登录 */
export async function logout(options?: { [key: string]: any }) {
  return umiRequest<Record<string, any>>('/api/core/auth/user/logout', {
    method: 'POST',
    headers: authHeaders(),
    ...(options || {}),
  });
}

/** 修改密码 */
export async function changePwd(data: Record<string, any>, options?: { [key: string]: any }) {
  return umiRequest<Record<string, any>>('/api/core/auth/user/changepwd', {
    method: 'POST',
    data,
    headers: authHeaders(),
    ...(options || {}),
  });
}
