import React from "react";

const ParamMonitor = ({ params }) => {
  return (
    <ul>
      {params.map(p => (
        <li key={p.path}>{p.path}</li>
      ))}
    </ul>
  );
};

export default ParamMonitor;
