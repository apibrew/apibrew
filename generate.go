package main

// Run code generation with: go generate bitbucket.org/supermoneygames/...
// gRPC documentation: https://grpc.io/
//go:generate rm -rf stub
//go:generate mkdir stub
//go:generate mkdir stub/model

//go:generate sh -c "cd proto; ls model/*.proto | xargs -I {} protoc --proto_path=. --go_out=../stub --go_opt=paths=source_relative --go-grpc_out=../stub --go-grpc_opt=paths=source_relative {}"

//go:generate sh -c "cd proto; ls *.proto | xargs -I {} protoc --proto_path=. --go_out=../stub --go_opt=paths=source_relative --go-grpc_out=../stub --go-grpc_opt=paths=source_relative {}"
