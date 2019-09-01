import { useState, useCallback } from "react";

const useParameters = () => {
  let [allParams, setAllParams] = useState([]);
  let [checkedParams, setCheckedParams] = useState([]);
  let updateParameters = () =>
    window.backend.Modules.GetAllParamInfo().then(p => {
      console.log("SDFDSF" + p);
      setAllParams(p);
    });
  updateParameters();
  return {
    allParams,
    setAllParams,
    checkedParams,
    setCheckedParams,
    updateParameters
  };
};

export default useParameters;
