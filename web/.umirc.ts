import { defineConfig } from '@umijs/max';
import proxy from './config/proxy';

export default defineConfig({
  antd: {
    // compact: true,
    configProvider: {
      theme: {
        "token": {
          "colorPrimary": "#0ec7a7",
          "colorInfo": "#0ec7a7",
          "colorBgLayout": "#ededed",
          "colorBorder": "#c1c1c1",
          "colorBorderSecondary": "#dcdcdc"
        },
        "components": {
          "Layout": {
            "bodyBg": "#ededed"
          }
        }
      }
    }
  },
  access: {},
  model: {},
  base: "/",
  manifest: {
    basePath: process.env.WEB_BASE !== undefined && process.env.WEB_BASE !== "" ? process.env.WEB_BASE : "",
  },
  publicPath: (process.env.WEB_BASE !== undefined && process.env.WEB_BASE !== "" ? process.env.WEB_BASE + '/' : '/'),
  locale: {
    // 默认使用 src/locales/zh-CN.ts 作为多语言文件
    default: 'zh-CN',
    baseSeparator: '-',
  },
  initialState: {},
  history: { type: 'hash' },
  request: {},
  dva: {},
  mfsu: false,
  layout: {
    title: '炘智科技'
  },
  define: { 'process.env.WEB_BASE': process.env.WEB_BASE },
  proxy: proxy['test'],
  alias: {
  },
  routes: [
    {
      path: '/',
      redirect: '/dashboard',
    },
    {
      name: 'Block',
      path: '/block',
      component: './block',
      layout: false
    },
    {
      name: '首页',
      path: '/home',
      component: './Home',
      layout: false
    },
    {
      path: '/user',
      layout: false,
      routes: [
        {
          path: '/user/login',
          layout: false,
          name: 'login',
          component: './user/Login',
        },
      ],
    },
    {
      path: '/dashboard',
      component: "dashboard",
      menuRender: false,
    },
    {
      path:'/demo/resizable/panel',
      component:"ResizablePanel",
      menuRender: false,
    },
    {
      path: '/form',
      layout: false,
      routes: [
        {
          name: '表单设计器',
          icon: 'smile',
          path: '/form/designable/:formID',
          component: './form/designable',
          layout: false,
        },
        {
          icon: 'smile',
          path: '/form/preview/:formID',
          component: './form/preview',
        },
      ]
    },
    {
      path: '/curd',
      routes: [
        {
          icon: 'smile',
          path: '/curd/page/view',
          component: './curd/CurdPage',
        },
        {
          icon: 'smile',
          exact: false,
          path: '/curd/page/manager/:name',
          component: './curd/CurdPage',
        },
        {
          icon: 'smile',
          exact: true,
          path: '/curd/page/edit',
          component: './curd/CurdPage/Edit',
        },
      ]
    },
    {
      path: 'bpm',
      layout: false,
      routes: [
        {
          name: '流程设计',
          icon: 'smile',
          path: '/bpm/designer/:processID',
          component: './bpm/DesignerPage',

        },
      ]
    },
    {
      path: '/cell',
      routes: [
        {
          icon: 'smile',
          exact: true,
          path: '/cell/edit',
          component: './cell/edit',
        },
      ]
    },
  ],
  npmClient: 'pnpm',
  jsMinifierOptions: {
    target: ['chrome80', 'es2020']
  },
  scripts: [
    'https://api.map.baidu.com/api?v=3.0&ak=H5pKkiqL3XGwDyNrGU3IBiQQvdj8RcOL',
    { src: (process.env.WEB_BASE !== undefined && process.env.WEB_BASE !== "" ? process.env.WEB_BASE : "") + `/js/editor/loader.js` }
  ],
});

