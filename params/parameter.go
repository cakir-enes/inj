package params

// type ParamType int8

// const (
// 	INT ParamType = 0
// 	FLT
// 	STR
// 	ARR
// )

type ParamInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}
