package params

import (
	"context"
	pb "gogui/injector"
)

type Accessor interface {
	Name() string
	Get(paramName string) (string, error)
	MultiGet(ctx context.Context, paramNames []string) (<-chan []*pb.PathAndValue, error)
	MultiSet(params map[string]string) error
	GetAllParamInfo() (*[]ParamInfo, error)
	GetAllPossibleEnums() (map[string][]string, error)
}
