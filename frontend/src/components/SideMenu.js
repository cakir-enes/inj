import React, { useState } from "react";
import { Layout, Menu, Breadcrumb, Icon } from "antd";

const { Header, Content, Footer, Sider } = Layout;
const { SubMenu } = Menu;

export const PAGE = {
  SELECT: "1",
  MONITOR: "2"
};

const SideMenu = ({ setActivePage }) => {
  let [collapsed, setCollapsed] = useState(true);
  return (
    <Sider collapsible collapsed={collapsed} onCollapse={c => setCollapsed(c)}>
      <div className="logo" />
      <Menu
        theme="dark"
        defaultSelectedKeys={["1"]}
        mode="inline"
        onSelect={p => setActivePage(p.key)}
      >
        <Menu.Item key={PAGE.SELECT}>
          <Icon type="check" />
          <span>Select Parameters</span>
        </Menu.Item>
        <Menu.Item key={PAGE.MONITOR}>
          <Icon type="eye" />
          <span>Monitor Selected Parameters</span>
        </Menu.Item>
      </Menu>
    </Sider>
  );
};

export default SideMenu;
