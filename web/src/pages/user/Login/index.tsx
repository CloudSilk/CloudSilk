import React from 'react';
import { history, useModel } from '@umijs/max';
import {useIntl, FormattedMessage} from 'react-intl'
import { pushWithRedirect } from '@swiftease/atali-pkg';

import { LoginComponent } from '@swiftease/atali-components';
import { defaultService } from '@swiftease/atali-form'



const Login: React.FC = () => {
  const { initialState, setInitialState } = useModel('@@initialState');

  // const intl = useIntl();

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();
    if (userInfo) {
      setInitialState({
        ...initialState,
        currentUser: userInfo,
      });
    }
  };
  return <LoginComponent showLogo={true} logo={(process.env.WEB_BASE!==undefined && process.env.WEB_BASE!=="" ? process.env.WEB_BASE : "") +'/logo-v.svg'} name={"智能工厂"} login={(params) => {
    return defaultService.login(params)
  }} formatMessage={(id: string, defaultMessage: string) => {
    return defaultMessage
    // return intl.formatMessage({ id: id, defaultMessage: defaultMessage })
  }}
    formattedMessage={(id: string, defaultMessage: string) => {
      return <FormattedMessage
        id={id}
        defaultMessage={defaultMessage}
      />
    }
    }
    fetchUserInfo={fetchUserInfo} redirect={() => {
      if (!history) return;
      setTimeout(() => {
        pushWithRedirect(history);
      }, 10);
    }}></LoginComponent>
};

export default Login;
