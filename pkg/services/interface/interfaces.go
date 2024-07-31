package interfaces

import (
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/enities"
	pb "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/pb"
)

type UserServiceInter interface {
	SignupService(userpb *pb.UserCreate) (*enities.User, error)
	LoginService(userpb *pb.LoginRequest) (*enities.User, error)
	FindAllUserService(p *pb.FetchAll) ([]*enities.User, error)
	EditUserService(p *pb.UserCreate) (*pb.UserCreate, error)
	FindUserByEmailService(p *pb.Email) (*pb.UserDetail, error)
	FindUserByIDService(p *pb.ID) (*pb.UserDetail, error)
	DeleteUserService(p *pb.ID) (*pb.CommonResponse, error)
}
