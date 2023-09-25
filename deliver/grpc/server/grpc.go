package server

import (
	"google.golang.org/grpc"

	pb "github/kunhou/simple-backend/deliver/grpc/proto/setting"
	"github/kunhou/simple-backend/deliver/grpc/router"
)

func NewGRPCServer(debug bool, settingRouter *router.SettingRouter) *grpc.Server {
	s := grpc.NewServer()

	pb.RegisterSettingServiceServer(s, settingRouter)

	return s
}
