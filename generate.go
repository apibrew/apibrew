package main

// Run code generation with: go generate bitbucket.org/supermoneygames/...
// gRPC documentation: https://grpc.io/
//go:generate rm -rf stub
//go:generate rm -rf model
//go:generate rm -rf grpc/stub
//go:generate mkdir model
//go:generate mkdir grpc/stub

//go:generate sh -c "cd proto; ls model/*.proto | xargs -I {} protoc --proto_path=. --go_out=../ --go_opt=paths=source_relative --go-grpc_out=../ --go-grpc_opt=paths=source_relative {}"

//go:generate sh -c "cd proto; ls *.proto | xargs -I {} protoc --proto_path=. --go_out=../grpc/stub --go_opt=paths=source_relative --go-grpc_out=../grpc/stub --go-grpc_opt=paths=source_relative {}"
