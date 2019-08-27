package main
//go:generate sh -c protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/*.proto
