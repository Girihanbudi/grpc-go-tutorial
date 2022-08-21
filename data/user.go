package data

import (
	"github.com/girihanbudi/grpc-tutorial/protos/user"
	proto "github.com/girihanbudi/grpc-tutorial/protos/user"
)

var superUserId = "c9bdcbcf-7955-4492-9c4e-225f1131ea55"
var superAdmin = proto.User{
	Id:       &superUserId,
	Role:     user.UserRoles_ADMIN.Enum(),
	Name:     "Giri Hanbudi",
	Email:    "ghanbudi@gmail.com",
	Password: "superpassword",
	Gender:   user.UserGender_MALE,
}

var UsersData = proto.UserList{
	List: []*proto.User{
		&superAdmin,
	},
}
