package params

import (
	"context"
	pb "gogui/injector"
	"io"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// GrpcAccessor is accessor impl for GRPC modules
type GrpcAccessor struct {
	name    string
	address string
	client  pb.InjectorClient
}

// New returns grpc acc. connected to given addr "localhost:6060", err if can't connect
func New(address string, name string) (*GrpcAccessor, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewInjectorClient(conn)
	return &GrpcAccessor{address, name, client}, nil
}

// Name returns connected module identifier
func (acc *GrpcAccessor) Name() string {
	return acc.name
}

// Get returns value of given parameter, i.e "Ship.Fuel.Remaning -> 3.2"
func (acc *GrpcAccessor) Get(paramName string) (string, error) {
	stream, err := acc.client.MultiGet(context.Background(), &pb.MultiGetReq{ParamPaths: []string{paramName}})
	if err != nil {
		println(err)
	}
	val, err := stream.Recv()
	if err == io.EOF {
		log.Printf("%s: Grpc Client Closed Connection\n", acc.name)
	}
	if err != nil {
		return "", err
	}
	return val.GetPairs()[0].GetValue(), nil
}

// MultiGet returns read channel for new values of given parameter paths.
func (acc *GrpcAccessor) MultiGet(ctx context.Context, paramNames []string) (<-chan []*pb.PathAndValue, <-chan error) {
	stream, err := acc.client.MultiGet(context.Background(), &pb.MultiGetReq{ParamPaths: paramNames})
	if err != nil {
		println(err)
	}
	valStream := make(chan []*pb.PathAndValue)
	errStream := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				stream.CloseSend()
				return
			default:
				vals, err := stream.Recv()
				if err == io.EOF {
					log.Printf("%s: Grpc Client Closed Connection\n", acc.name)
					break
				}
				if err != nil {
					errStream <- err
					break
				}
				valStream <- vals.Pairs
			}
		}
	}()
	return valStream, errStream
}

// MultiSet tries setting given params to given values.
func (acc *GrpcAccessor) MultiSet(params map[string]string) error {
	_, err := acc.client.MultiSet(context.Background(), &pb.MultiSetReq{PathToVal: params})
	if err != nil {
		return err
	}
	return nil
}

// GetAllParamInfo returns all the parameter paths, types, values from connected module.
func (acc *GrpcAccessor) GetAllParamInfo() ([]ParamInfo, error) {
	resp, err := acc.client.GetAllParameterInfos(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	protoParams := resp.GetParamInfos()
	var params = make([]ParamInfo, len(protoParams))
	for i, p := range protoParams {
		params[i] = ParamInfo{Name: p.GetPath(), Type: p.GetType().String(), Value: p.GetValue()}
	}
	return params, nil
}

// GetAllPossibleEnums returns map of all enum types to values of connected module
func (acc *GrpcAccessor) GetAllPossibleEnums() (map[string][]string, error) {
	resp, err := acc.client.GetAllEnumValues(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	enums := make(map[string][]string)
	for k, v := range resp.GetValues() {
		enums[k] = v.GetValues()
	}
	return enums, nil
}
