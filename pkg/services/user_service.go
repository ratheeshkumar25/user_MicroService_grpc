package services

import (
	"errors"
	"fmt"

	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/enities"
	pb "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/pb"
	interfaces "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/repo/interface"
	inter "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/services/interface"
)

type UserService struct {
	repo interfaces.UserRepoInter
}

func NewUserService(repo interfaces.UserRepoInter) inter.UserServiceInter {
	return &UserService{
		repo: repo,
	}
}

// SignupService implements interfaces.UserServiceInter.
func (u *UserService) SignupService(userpb *pb.UserCreate) (*enities.User, error) {
	user := &enities.User{
		UserName: userpb.Username,
		Email:    userpb.Email,
		Password: userpb.Password,
		Role:     "user",
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// LoginService implements interfaces.UserServiceInter.
func (u *UserService) LoginService(userpb *pb.LoginRequest) (*enities.User, error) {
	user, err := u.repo.FindUserByEmail(userpb.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != userpb.Password {
		return nil, errors.New("invalid password")
	}
	if user.Role != userpb.Role {
		return nil, errors.New("invalid role")
	}
	return user, nil

}

// FindAllUserService implements interfaces.UserServiceInter.
func (u *UserService) FindAllUserService(p *pb.FetchAll) ([]*enities.User, error) {
	users, err := u.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// EditUserService implements interfaces.UserServiceInter.
func (u *UserService) EditUserService(p *pb.UserCreate) (*pb.UserCreate, error) {
	var usermodel enities.User
	usermodel.UserName = p.Username
	usermodel.Email = p.Email
	usermodel.Password = p.Password
	err := u.repo.UpdateUser(&usermodel)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// FindUserByIDService implements interfaces.UserServiceInter.
func (u *UserService) FindUserByIDService(p *pb.ID) (*pb.UserDetail, error) {
	user, err := u.repo.FindUserByID(uint(p.Id))
	if err != nil {
		fmt.Println("psoioi")
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.UserDetail{
		Id:       uint32(user.ID),
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

// FindUserByEmailService implements interfaces.UserServiceInter.
func (u *UserService) FindUserByEmailService(p *pb.Email) (*pb.UserDetail, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	return &pb.UserDetail{
		Id:       uint32(user.ID),
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

// DeleteUserService implements interfaces.UserServiceInter.
func (u *UserService) DeleteUserService(p *pb.ID) (*pb.CommonResponse, error) {
	err := u.repo.DeleteUser(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "user deleted",
	}, nil
}
