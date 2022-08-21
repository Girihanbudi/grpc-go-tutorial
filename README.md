# Simple grpc example
This project is a simple go template to learn grpc mechanism. You can use this repo as your starting point for making grpc based service.

### How to setup
Install minGW to run make file
> https://stackoverflow.com/questions/2532234/how-to-run-a-makefile-in-windows

Install Protocol Buffer from release page
> https://developers.google.com/protocol-buffers/docs/downloads

Install gRPC setup
> https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

Install google RPC as go dependency
```go get google.golang.org/grpc```

### How to generate protobuf in go format
In Makefile code, there is an example protobuff call command to generate pb file. Create your own custom .proto file in protos folder and add protobuff call command for exact proto file

### How to do reflection to inspect method on grpc active service
Get all grpc reflection
```grpcurl --plaintext localhost:8080 list```
Get all grpc reflection on group
```grpcurl --plaintext localhost:8080 list <group name>```
Get describtion of grpc method
```grpcurl --plaintext localhost:8080 describe User.GetUser```
Do request (on windows double quote json need to be escaped)
```grpcurl --plaintext -d '{\"Email\":\"ghanbudi@gmail.com\"}' localhost:8080 User.GetUser```

### Test the code
You can do some automation test on GRPC method using bellow syntax
```go test ./... -v -coverpkg=./...```