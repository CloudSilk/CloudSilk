import React, { useCallback, useState } from 'react';
import { LogoutOutlined, SettingOutlined, UserOutlined } from '@ant-design/icons';
import { Avatar, Button, Form, Input, Menu, message, Modal, Spin } from 'antd';
import { history, request, useModel } from '@umijs/max';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';
import type { MenuInfo } from 'rc-menu/lib/interface';
import { getToken,  removeToken } from '@swiftease/atali-pkg';
import { replaceTakeRedirect } from '@swiftease/atali-pkg';
import { defaultCache } from '@swiftease/atali-form';
export type GlobalHeaderRightProps = {
  menu?: boolean;
};

async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/core/auth/user/logout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      authorization: 'Bearer ' + getToken(),
    },
    ...(options || {}),
  });
}

async function changePwd(data: any, options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/core/auth/user/changepwd', {
    method: 'POST',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      authorization: 'Bearer ' + getToken(),
    },
    ...(options || {}),
  });
}

/**
 * 退出登录，并且将当前的 url 保存
 */
const loginOut = async () => {
  await logout();
  removeToken();
  replaceTakeRedirect(history, '/user/login', '/user/login')
};

const AvatarDropdown: React.FC<GlobalHeaderRightProps> = ({ menu }) => {
  const { initialState, setInitialState } = useModel('@@initialState');

  const onMenuClick = useCallback(
    (event: MenuInfo) => {
      const { key } = event;
      if (key === 'logout' && initialState) {
        setInitialState({ ...initialState, currentUser: undefined });
        loginOut();
        return;
      }
      if (key === "changepwd") {
        showModal()
      }
      if (key === "clearLocalStorage") {
        defaultCache.clearLocalStorage()
        location.reload()
      }
      if (key === "enableLocalStorage") {
        defaultCache.enableCache()
        defaultCache.clearLocalStorage()
        location.reload()
      }
      if (key === "disableLocalStorage") {
        defaultCache.disableCache()
        defaultCache.clearLocalStorage()
        location.reload()
      }
    },
    [initialState, setInitialState],
  );

  const [isModalVisible, setIsModalVisible] = useState(false);
  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = () => {
    setIsModalVisible(false);
  };
  const handleCancel = () => {
    setIsModalVisible(false);
  };
  const onFinish = (values: any) => {
    if (values.oldPwd === values.newPwd) {
      message.success('新密码和旧密码一样!')
      return
    }

    if (values.newConfirmPwd !== values.newPwd) {
      message.success('新密码和确认密码不一样!')
      return
    }

    changePwd(values).then(resp => {
      if (resp.code === 20000) {
        handleCancel()
        message.success('密码修改成功!')
      } else {
        message.success('密码修改失败!')
      }
    })
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };
  const loading = (
    <span className={`${styles.action} ${styles.account}`}>
      <Spin
        size="small"
        style={{
          marginLeft: 8,
          marginRight: 8,
        }}
      />
    </span>
  );

  if (!initialState) {
    return loading;
  }

  const { currentUser } = initialState;

  if (!currentUser || !currentUser.nickname) {
    return loading;
  }
  const cacheStatus = defaultCache.isEnableCache()

  const menuHeaderDropdown = (
    <Menu className={styles.menu} selectedKeys={[]} onClick={onMenuClick}>
      {menu && (
        <Menu.Item key="center">
          <UserOutlined />
          个人中心
        </Menu.Item>
      )}
      {menu && cacheStatus !== true && (
        <Menu.Item key="enableLocalStorage">
          <SettingOutlined />
          启用缓存
        </Menu.Item>
      )}
      {menu && cacheStatus === true && (
        <Menu.Item key="disableLocalStorage">
          <SettingOutlined />
          禁用缓存
        </Menu.Item>
      )}
      {menu && cacheStatus === true && (
        <Menu.Item key="clearLocalStorage">
          <SettingOutlined />
          清空缓存
        </Menu.Item>
      )}
      {menu && <Menu.Divider />}
      {menu && (
        <Menu.Item key="changepwd">
          <SettingOutlined />
          修改密码
        </Menu.Item>
      )}
      {menu && <Menu.Divider />}

      <Menu.Item key="logout">
        <LogoutOutlined />
        退出登录
      </Menu.Item>
    </Menu>
  );
  const avatar=!currentUser.avatar || currentUser.avatar===""?"/icon-32x32.png":currentUser.avatar
  return (
    <>
      <HeaderDropdown overlay={menuHeaderDropdown}>
        <span className={`${styles.action} ${styles.account}`}>
          <Avatar size="small" className={styles.avatar} src={avatar} alt="avatar" />
          <span className={`${styles.name} anticon`} style={{color:'white'}}>{currentUser.nickname}</span>
        </span>
      </HeaderDropdown>
      <Modal title="修改密码" footer={null} open={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
        <Form
          name="basic"
          labelCol={{ span: 8 }}
          wrapperCol={{ span: 16 }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item
            label="旧密码"
            name="oldPwd"
            rules={[{ required: true, message: '请输入旧密码!' }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item
            label="新密码"
            name="newPwd"
            rules={[{ required: true, message: '请输入新密码!' }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item
            label="确认密码"
            name="newConfirmPwd"
            rules={[{ required: true, message: '请输入确认密码!' }]}
          >
            <Input.Password />
          </Form.Item>
          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Button type="primary" htmlType="submit">
              保存
            </Button>
            <Button onClick={handleCancel}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default AvatarDropdown;
