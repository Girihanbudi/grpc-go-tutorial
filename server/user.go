package server

import (
	"context"
	"errors"

	"github.com/girihanbudi/grpc-tutorial/data"
	proto "github.com/girihanbudi/grpc-tutorial/protos/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	log hclog.Logger
	proto.UnimplementedUsersServer
}

func NewUserService(log hclog.Logger) *UserService {
	return &UserService{
		log: log,
	}
}

func (u *UserService) Register(ctx context.Context, in *proto.User) (*emptypb.Empty, error) {
	u.log.Info("Handle Register User", "Data", in)
	emptyRes := new(empty.Empty)

	for _, user := range data.UsersData.List {
		if user.Email == in.Email {
			u.log.Error("Handle Register User", "User already exist")
			return emptyRes, errors.New("user already exist")
		}
	}

	newID, err := uuid.NewRandom()
	if err != nil {
		u.log.Error("Handle Register User", "Failed to generate uuid")
		return emptyRes, errors.New("failed to generate uuid")
	}
	newStringID := newID.String()
	in.Id = &newStringID
	in.Role = proto.UserRoles_USER.Enum()

	data.UsersData.List = append(data.UsersData.List, in)

	return emptyRes, nil
}

func (u *UserService) List(ctx context.Context, in *emptypb.Empty) (*proto.UserList, error) {
	u.log.Info("Handle Get User List")
	return &data.UsersData, nil
}

func (u *UserService) UserDetail(ctx context.Context, in *proto.GetUserDetail) (*proto.User, error) {
	u.log.Info("Handle Get User Detail", "Data", in)

	for _, user := range data.UsersData.List {
		if user.Email == in.Email {
			return user, nil
		}
	}

	u.log.Error("Handle Get User Detail", "User is not exist")
	return &proto.User{}, errors.New("user is not exist")
}
