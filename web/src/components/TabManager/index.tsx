import React, { createContext, useEffect, useState } from 'react';
import { Button, Dropdown, Menu, Tabs } from 'antd';
import { history, useLocation, useModel } from '@umijs/max';
import queryString from 'query-string';

import './index.less';
import { RecursiveCall } from '@swiftease/atali-pkg';
import { DownOutlined } from '@ant-design/icons';

const TabContext = createContext({});

function getDefultMenu(menus: any, path: string, selectedTab: any[]) {
    RecursiveCall(menus, (menu: any) => {
      if (menu.parentID && menu.path === path) {
        console.log('currentMenu: ', menu);
        selectedTab.push(menu);
        return;
      }
    });
}

const TabManager = ({ children }: any) => {
  const [activeKey, setActiveKey] = useState('0');
  const {tabs, setTabs} = useModel('global');
  const {initialState}: any  = useModel('@@initialState');
  const menus = initialState?.currentUser?.menus || [];
  const location = useLocation();
  const query = queryString.parse(location.search);
  const pageName = query?.pageName ?? '';
  
  useEffect(() => {
    if(!tabs.length && pageName) {
      const selectedTab:any[] = [];
      getDefultMenu(menus, (location.pathname + location.search), selectedTab);
      if(selectedTab.length) {
        tabs.push(selectedTab[0]);
        setTabs(tabs);
        setActiveKey(selectedTab[0].id);
      }
    } else {
      const tab = tabs.find(tab => tab.path === (location.pathname + location.search));
      if(tab) {
        setActiveKey(tab.id);
      }
    }
    
  }, [pageName]);
  const goTopage = (path: string) => {
    history.push(path);
  }

  const onEdit = (targetKey: any, action: any) => {
    if (action === 'add') {
      const newKey = `newTab${Date.now()}`;
      setActiveKey(newKey);
    } else if (action === 'remove') {
      if(tabs.length <= 1) {
        return;
      }
      if(activeKey === targetKey) {
        setActiveKey(tabs[tabs.length - 2].key);
        goTopage(tabs[tabs.length - 2].path);
      }
      setTabs(tabs.filter(pane => pane.id !== targetKey));
      
    }
  };

  const moreActionClick = ({ item, key, keyPath, domEvent }: any) => {

    if(key === '1') {
      tabs.length = 1;
      setTabs(tabs);
      setActiveKey(tabs[0].id);
      goTopage(tabs[0].path);
    }
  }

  return (
    <TabContext.Provider value={{ activeKey, setActiveKey }}>
      <div className='t-m-container'>
        <Tabs
        type="editable-card"
        items={tabs.map(tab => ({
          label: tab.name,
          key: tab.id,
          
          closable: tabs.length <= 1 ? false : true,
        }))}
        
        onEdit={onEdit}
        activeKey={activeKey}
        onChange={(tabId) => {
          setActiveKey(tabId);
          const tab = tabs.find(item => item.id === tabId);
          if(tab) {
            goTopage(tab.path);
          }
        }}
        hideAdd={true}
        tabBarExtraContent={
          <Dropdown
            overlay={
              <Menu onClick={moreActionClick}>
                <Menu.Item key="1">关闭全部</Menu.Item>
              </Menu>
            }
          >
            <Button type="link">
              更多操作 <DownOutlined />
            </Button>
          </Dropdown>}
      >

      </Tabs>
      <div style={{marginTop:-10}}>
      {children}
      </div>
      </div>
      
    </TabContext.Provider>
  );
};

export { TabManager, TabContext };