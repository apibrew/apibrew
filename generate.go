package main

// Run code generation with: go generate bitbucket.org/supermoneygames/...
// gRPC documentation: https://grpc.io/
//go:generate rm -rf model
//go:generate mkdir model

//go:generate sh -c "cd proto; ls model/*.proto | xargs -I {} protoc --proto_path=. --go_out=../ --go_opt=paths=source_relative --go-grpc_out=../ --go-grpc_opt=paths=source_relative {}"

//go:generate sh -c "cd proto; buf generate"
//go:generate sh -c "mv stub server/"
