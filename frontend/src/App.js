import React, { useState, useEffect } from "react";
import logo from "./logo.png";
import { Layout } from "antd";
import "antd/dist/antd.css";
import "./App.css";
import HelloWorld from "./components/HelloWorld";
import ParamTree from "./components/ParamTree";
import useParameters from "./hooks";
import Sider from "antd/lib/layout/Sider";
import SiderDemo, { PAGE } from "./components/SideMenu";

// function useParameters() {
//   let [params, setParams] = useState([]);
//   params = api.parameters;
// }
const api = {
  // get parameters() {
  //   let c = await window.backend.basic()
  // }
};

function App() {
  // let paramApi = useParameters();
  let [allParams, setAllParams] = useState([]);
  let [checkedParams, setCheckedParams] = useState([]);
  let [activePage, setActivePage] = useState(PAGE.SELECT);
  let updateParameters = () => {
    window.backend.Modules.GetAllParamInfo().then(p => {
      // console.log("B" + p);
      setAllParams(p);
    });
  };
  useEffect(() => updateParameters(), []);
  let paramApi = {
    allParams,
    setAllParams,
    checkedParams,
    setCheckedParams
    // updateParameters
  };
  let renderActivePage = p => {
    switch (p) {
      case PAGE.MONITOR:
        return <h3>HELHELELO</h3>;
      case PAGE.SELECT:
        return (
          <ParamTree
            params={paramApi.allParams}
            checked={paramApi.checkedParams}
            setChecked={paramApi.setCheckedParams}
          />
        );
      default:
        return <h1>NO</h1>;
    }
  };
  return (
    <div id="app" className="App">
      <Layout style={{ minHeight: "100vh" }}>
        <SiderDemo setActivePage={p => setActivePage(p)} />
        {renderActivePage(activePage)}
      </Layout>
    </div>
  );
}

export default App;
