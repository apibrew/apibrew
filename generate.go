package main

// Run code generation with: go generate bitbucket.org/supermoneygames/...
// gRPC documentation: https://grpc.io/
//go:generate rm -rf stub
//go:generate mkdir -p stub
//go:generate protoc --proto_path=proto --go_out=stub --go_opt=paths=source_relative --go-grpc_out=stub --go-grpc_opt=paths=source_relative resource.proto
//go:generate protoc --proto_path=proto --go_out=stub --go_opt=paths=source_relative --go-grpc_out=stub --go-grpc_opt=paths=source_relative system.proto
//go:generate protoc --proto_path=proto --go_out=stub --go_opt=paths=source_relative --go-grpc_out=stub --go-grpc_opt=paths=source_relative generic.proto
