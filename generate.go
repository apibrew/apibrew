package main

// Run code generation with: go generate bitbucket.org/supermoneygames/...
// gRPC documentation: https://grpc.io/
//go:generate rm -rf stub
//go:generate rm -rf stub-ts
//go:generate mkdir stub
//go:generate mkdir stub/model
//go:generate mkdir stub-ts

//go:generate sh -c "ls proto/model | xargs -I {} protoc --ts_out=import_style=commonjs:stub-ts --proto_path=proto --go_out=stub --go_opt=paths=source_relative --go-grpc_out=stub --go-grpc_opt=paths=source_relative model/{}"

//go:generate sh -c "ls proto | xargs -I {} protoc --ts_out=import_style=commonjs:stub-ts --proto_path=proto --go_out=stub --go_opt=paths=source_relative --go-grpc_out=stub --go-grpc_opt=paths=source_relative {}"
