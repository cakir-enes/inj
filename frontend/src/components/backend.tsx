export interface Parameter {
  name: string;
  value: string;
  type: "ARR" | "STR" | "INT" | "FLT" | "OBJ";
}
declare global {
  interface Window {
    Modules: {
      GetAllParamInfo: () => Array<Parameter>;
    };
  }
}
