import { Settings as LayoutSettings } from '@ant-design/pro-layout';
import { TopNavHeader } from '@ant-design/pro-components';
import { clearMenuItem } from '@ant-design/pro-layout/lib/utils/utils';
import { history, Link, useModel } from '@umijs/max';
import RightContent from '@/components/RightContent';
import { GetDetailResponse, User, newResponseInterceptor, getToken, replaceTakeRedirect, Menu as AtaliMenu, RecursiveCall } from '@swiftease/atali-pkg';
import umiRequest from 'umi-request';
import { MyIcon } from '@swiftease/atali-form';
import { notification } from 'antd';
import './app.less'
const loginPath = '/user/login';
// process.env.NODE_ENV = 'production';
import * as monaco from "monaco-editor";
import { loader } from "@monaco-editor/react";
loader.config({ monaco });

umiRequest.interceptors.response.use(newResponseInterceptor(() => {
  replaceTakeRedirect(history, '/user/login', '/user/login')
},
  (description: string, message: string) => {
    notification.error({
      description: description, message: message
    })
  }))

async function queryCurrentUser(options?: { [key: string]: any }) {
  return umiRequest<GetDetailResponse<User.UserProfile>>(
    '/api/core/auth/user/profile',
    {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        authorization: 'Bearer ' + getToken(),
      },
      ...(options || {}),
    },
  );
}

/** 获取用户信息比较慢的时候会展示一个 loading */
// export const initialStateConfig = {
//   loading: <PageLoading />,
// };

export interface AppState {
  settings?: Partial<LayoutSettings>;
  currentUser?: User.UserProfile;
  fetchUserInfo?: () => Promise<User.UserProfile | undefined>;
  currentMenus?: AtaliMenu[]
}
/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<AppState> {
  const fetchUserInfo = async () => {
    try {
      const currentUser = await queryCurrentUser();
      console.log('currentUser: ', currentUser);
      return currentUser.data;
    } catch (error) {
      replaceTakeRedirect(history, loginPath, loginPath);
    }
    return undefined;
  };
  // 如果是登录页面，不执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
      settings: {}
    };
  }
  return {
    fetchUserInfo,
    settings: {}
  }
};

function getDefaultMenuPath(menus: any, menuID: string, defaultPath: string) {
  let p = defaultPath;
  menus.forEach((m: any) => {
    if (m.id == menuID) {
      RecursiveCall(m.children, (menu: any) => {
        if (menu.defaultMenu) {
          p = menu.path;
        }
      })
    }
  });
  return p;
}

export const layout = ({
  initialState, setInitialState
}: {
  initialState: AppState;
  setInitialState: any
}) => {

  const { tabs, setTabs } = useModel('global');
  return {
    logo: (process.env.WEB_BASE !== undefined && process.env.WEB_BASE !== "" ? process.env.WEB_BASE : "") + "/icon-32x32.png",
    layout: 'mix',
    siderWidth: 230,
    contentWidth: 'Fluid',
    fixedHeader: true,
    fixSiderbar: true,
    splitMenus: true,
    rightContentRender: () => <RightContent />,
    menuFooterRender: (props) => {
      if (props?.collapsed) return undefined;
      return (
        <div
          style={{
            textAlign: 'center',
            paddingBlockStart: 12,
          }}
        >
          <div>© 2024 上海炘智科技有限公司</div>
        </div>
      );
    },
    // footerRender: () => <Footer />,
    disableContentMargin: true,
    waterMarkProps: {
      content: initialState?.currentUser?.nickname,
    },
    breadcrumbRender: /* (routes) => [
      {
        path: '/',
        breadcrumbName: '主页',
      },
      ...(routes || []),
    ] */ false,
    onPageChange: () => {
      const { currentUser } = initialState;
      const { location } = history;
      // 如果没有登录，重定向到 login
      if (!currentUser && location.pathname !== loginPath) {
        history.push(loginPath);
      }
    },
    token: {
      bgLayout: "#f3f4f6",
      header: {
        colorBgHeader: '#0ec7a7',
        colorBgMenuItemSelected: "#06A88D",
        colorHeaderTitle: '#fff',
        colorTextMenu: '#fff',
        colorTextMenuSecondary: '#dfdfdf',
        colorTextMenuSelected: '#fff',
        colorTextRightActionsItem: '#dfdfdf',
      },
      sider: {
        colorMenuBackground: '#DDE1EA',
        colorTextMenuSelected: '#0ec7a7',
        colorBgMenuItemSelected: '#e6fff7',
      },
      pageContainer: {
        paddingBlockPageContainerContent: 0,
        paddingInlinePageContainerContent: 0
      }
    },
    iconfontUrl: "//at.alicdn.com/t/font_2590742_729agyh0ndx.js",
    menuItemRender: (item, dom, props) => {
      // TODO 如何在这边点击的时候更新tab
      let path = item.path;
      if (item.parentID == "") {
        path = getDefaultMenuPath(initialState?.currentUser?.menus, item.id, item.path)
      }
      const handleClick = (item: any) => {
        if (item.parentID && item.path) {
          const exist = tabs.some((tab) => tab.id === item.id);
          if (!exist) {
            tabs.push({
              ...item,
              path: path,
            });
            setTabs(tabs);
          }

        } else {
          tabs.length = 0;
          setTabs(tabs);
        }
      }

      return <Link to={path} onClick={() => { handleClick(item) }}>
        <div className={dom.props.className}>
          <span onClick={item.onClick} className={props.baseClassName + "-item-icon " + props.hashId}>
            {<MyIcon type={(item.icon as string) || 'icon-daima'} />}
          </span>
          <span className={props.baseClassName + "-item-text " + props.hashId + " "}>{item.title}</span>
        </div>
      </Link>
    },
    subMenuItemRender: (item, dom, props) => {
      var iconEl = (props.collapsed && item.level <= 1) ? <div className={dom.props.className}>
        <span onClick={item.onClick} className={props.className}>
          {<MyIcon type={(item.icon as string) || 'icon-daima'} />}
        </span>
      </div> : <span onClick={item.onClick} className={props.className}>
        {<MyIcon style={{ marginRight: 5 }} type={(item.icon as string) || 'icon-daima'} />}
      </span>
      return <>{iconEl}
        <span className={props.className + " " + (props.collapsed ? 'ant-r-base-menu-vetical-item-text-has-icon' : 'ant-pro-ae-menu-inline-tem-text-as-icon')}>{item.title}</span>
      </>
    },
    ...initialState?.settings,
    menu: {
      locale: false,
      params: { currentMenus: initialState?.currentUser?.menus },
      request: (params, defaultMenuData) => {
        return params.currentMenus
      },
    },
    className: "my-app-layout",
    headerRender: (props: any) => {
      const menuData = []
      props.menuData?.forEach((value: any) => {
        menuData.push({ ...value, children: [] })
      })
      const clearMenuData = clearMenuItem(menuData || []);
      return <TopNavHeader
        mode="horizontal"
        onCollapse={props.onCollapse}
        {...props}
        menuData={clearMenuData}
        menuProps={{
          style: {
            //控制头部菜单居中显示
            width: "max-content",
            marginLeft: "auto",
            marginRight: "auto"
          }
        }}
      />
    }
  };
};