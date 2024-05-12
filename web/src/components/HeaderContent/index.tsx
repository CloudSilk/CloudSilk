// import { HeaderViewProps } from '@ant-design/pro-layout';
import { Menu } from 'antd';
import React from 'react';
import { Menu as AtaliMenu } from '@swiftease/atali-pkg';
import { MyIcon } from '@swiftease/atali-form';
import { History } from '@umijs/max';
interface HeaderContentProps {
  headerViewProps: any
  menus: AtaliMenu[]
  changeSystem: (menus: AtaliMenu[]) => void
  history: History
}

interface HeaderContentState {
  defaultSelectedKey: string
  menusID?: Map<string, AtaliMenu>;
  menusPath?: Map<string, AtaliMenu>;
}

export class HeaderContent extends React.Component<HeaderContentProps, HeaderContentState> {
  constructor(props: HeaderContentProps) {
    super(props)
    const menusID = new Map<string, AtaliMenu>();
    const menusPath = new Map<string, AtaliMenu>();
    this.props.menus.forEach(m => {
      menusID.set(m.id ?? '', m)
      const path = m.path.split("?")[0]
      menusPath.set(path, m)
    })
    this.state = {
      menusID: menusID,
      menusPath: menusPath,
      defaultSelectedKey: 'Dashboard'
    }
  }

  getMenus(system: string, currentPath: string) {
    const menus = this.props.menus
    if (menus) {
      for (let index = 0; index < menus.length; index++) {
        const element = menus[index];
        if (element.id === system) {
          element.children?.forEach((m: AtaliMenu) => {
            if (m.children?.length > 0) {
              m.children.forEach((child: AtaliMenu) => {
                if (child.defaultMenu) {
                  this.props.history.replace(child.path)
                }
              })
            }
          })
          return element.children
        }
      }
    }
    return [];
  }

  getMenuByPath(currentPath: string, menusPath?: Map<string, AtaliMenu>) {
    if (!menusPath) return undefined
    return menusPath.get(currentPath)
  }

  getSystem(id: string, menus: Map<string, AtaliMenu>): AtaliMenu | undefined {
    const menu = menus.get(id)
    if (!menu?.parentID || menu.parentID === '') return menu
    return this.getSystem(menu.parentID, menus)
  }

  render() {
    return <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['Dashboard']}>
      {this.props.menus.map(item => {
        return <Menu.Item key={item.name} onClick={() => {
          if (item.children?.length > 0) {
            this.props.changeSystem(this.getMenus(item.id ?? '', this.props.history.location.pathname))
          }
          else {
            this.props.history.replace(item.path)
            this.props.changeSystem([])
          }
        }} icon={<MyIcon type={item.icon} />}>
          {item.title}
        </Menu.Item>
      })}
    </Menu>
  }
}