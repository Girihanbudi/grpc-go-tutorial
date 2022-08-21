package server

import (
	"context"
	"testing"

	proto "github.com/girihanbudi/grpc-tutorial/protos/user"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = "localhost:8080"

// Test Register calls from gprc client
func TestRegister(t *testing.T) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("failed to connect to grpc server %s", err.Error())
	}

	client := proto.NewUsersClient(conn)

	if res, err := client.Register(context.Background(), &proto.User{
		Name:     "John Doe",
		Email:    "john-doe@gmail.com",
		Password: "johndoe123",
		Gender:   *proto.UserGender_MALE.Enum(),
	}); err != nil {
		t.Errorf("Failed to register new user %s", err.Error())
	} else {
		t.Logf("Success to call User.Register. Res : %s", res)
	}
}

// Test List calls from gprc client
func TestList(t *testing.T) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("failed to connect to grpc server %s", err.Error())
	}

	client := proto.NewUsersClient(conn)

	if res, err := client.List(context.Background(), new(empty.Empty)); err != nil {
		t.Errorf("Failed to get user list %s", err.Error())
	} else {
		t.Logf("Success to call User.List. Res : %s", res)
	}
}

// Test UserDetail calls from gprc client on inserted previous data
func TestUserDetailExist(t *testing.T) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("failed to connect to grpc server %s", err.Error())
	}

	client := proto.NewUsersClient(conn)

	if res, err := client.UserDetail(context.Background(), &proto.GetUserDetail{
		Email: "john-doe@gmail.com",
	}); err != nil {
		t.Errorf("Failed to get user detail %s", err.Error())
	} else {
		t.Logf("Success to call User.UserDetail. Res : %s", res)
	}
}

// Test UserDetail calls from gprc client on not exist data
func TestUserDetailNil(t *testing.T) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("failed to connect to grpc server %s", err.Error())
	}

	client := proto.NewUsersClient(conn)

	if _, err := client.UserDetail(context.Background(), &proto.GetUserDetail{
		Email: "jack@gmail.com",
	}); err == nil {
		t.Error("Grpc call should contain error")
	} else {
		t.Logf("Success to call User.UserDetail with not exist user. Res : %s", err)
	}
}
