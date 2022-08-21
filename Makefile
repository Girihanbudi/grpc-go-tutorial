.PHONY: protos

# run 'MinGW32-make protos' or 'make protos' to generate go protos rpc protocol
protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/user/user.proto