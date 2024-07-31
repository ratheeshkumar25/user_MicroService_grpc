package di

import (
	"github.com/ratheeshkumar/microservice_grpc_userV1/config"
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/db"
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/handler"
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/repo"
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/server"
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/services"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	userRepo := repo.NewUserRepo(db)
	userSvc := services.NewUserService(userRepo)
	userHandle := handler.NewUserHandler(userSvc)
	server.NewGrpcServer(userHandle)
}
