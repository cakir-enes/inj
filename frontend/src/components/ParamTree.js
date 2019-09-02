import React, { useState, useEffect } from "react";
import { Tree } from "antd";

const { TreeNode } = Tree;

function insert(node, path) {
  const pathParts = path.split(".");
  let current = node;

  pathParts.reduce((pre, curr) => {
    let children = current.children || [];
    let alreadyExists = false;
    let idx = children.findIndex(ch => ch.title === curr);
    alreadyExists = idx >= 0;

    if (!alreadyExists) {
      let node = {
        title: curr,
        key: pre + "." + curr,
        children: []
      };
      children.push(node);
      current = node;
    } else {
      current = children[idx];
    }
    return pre + "." + curr;
  });
}

const ParamTree = ({ params, checked, setChecked }) => {
  useEffect(() => console.log("CHECKED:" + checked), [checked]);
  let [selectedKeys, setSelectedkeys] = useState([]);
  useEffect(() => {
    console.log(selectedKeys);
  }, [selectedKeys]);
  let renderTreeNodes = data =>
    data.map(item => {
      if (item.children) {
        return (
          <TreeNode title={item.title} key={item.key} dataRef={item}>
            {renderTreeNodes(item.children)}
          </TreeNode>
        );
      }
      return <TreeNode key={item.key} {...item} />;
    });

  return (
    <Tree
      checkable
      //   autoExpandParent={expandParent}
      //   onExpand={e => {
      //     setExpanedKeys(e);
      //     setExpandParent(false);
      //   }}
      //   expandedKeys={expandedKeys}
      onSelect={s => setSelectedkeys(s)}
      selectedKeys={selectedKeys}
      onCheck={c => setChecked(c)}
      checkedKeys={checked}
      //   onSelect={s => setSelectedKeys(s)}
      //   selectedKeys={selectedKeys}
    >
      {renderTreeNodes(
        Object.keys(params)
          .map(p => {
            let root = { title: p, key: p, children: [] };
            params[p].forEach((n, i) => insert(root, n.name, i));
            return root;
          })
          .flat()
      )}
    </Tree>
  );
};

export default ParamTree;
