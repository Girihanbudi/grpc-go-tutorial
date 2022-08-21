package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/girihanbudi/grpc-tutorial/protos/user"
	"github.com/girihanbudi/grpc-tutorial/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	grpcServer := grpc.NewServer()
	userService := server.NewUserService(log)

	protos.RegisterUsersServer(grpcServer, userService)

	// For development only
	reflection.Register(grpcServer)

	targetPort := 8080
	log.Info("Service Starting Up", fmt.Sprintf("Listening on port %d", targetPort))
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", targetPort))
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(l)
}
