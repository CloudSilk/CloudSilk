import { useIntl } from 'react-intl';
import { DefaultFooter } from '@ant-design/pro-layout';
import './index.less';

export default () => {
  // const intl = useIntl();
  // console.log(intl)
  // const defaultMessage = intl.formatMessage({
  //   id: 'app.copyright.produced',
  //   defaultMessage: '上海炘智技术有限公司'
  // });

  return (
    <div style={{
      position: 'fixed', bottom: 0,width:'100%'
    }}>
      <DefaultFooter
        copyright={`2023 上海炘智技术有限公司`}
        links={[
          {
            key: '工站配置',
            title: "工站配置",
            href: '/aiot/#/aiot/station/config',
            blankTarget: true,
          },
        ]}
      /></div>
  );
};
