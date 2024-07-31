package interfaces

import "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/enities"

type UserRepoInter interface {
	CreateUser(user *enities.User) error
	FindUserByID(id uint) (*enities.User, error)
	UpdateUser(user *enities.User) error
	FindUserByEmail(email string) (*enities.User, error)
	FindAllUsers() ([]*enities.User, error)
	DeleteUser(id uint) error
}
