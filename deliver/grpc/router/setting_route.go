package router

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"

	pb "github/kunhou/simple-backend/deliver/grpc/proto/setting"
	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/usecase/setting"
)

type SettingRouter struct {
	pb.UnimplementedSettingServiceServer

	uc *setting.SettingUsecase
}

func NewSettingRouter(s *setting.SettingUsecase) *SettingRouter {
	return &SettingRouter{
		uc: s,
	}
}

func (s *SettingRouter) GetSettingByName(ctx context.Context, req *pb.GetSettingReq) (*pb.Setting, error) {
	res, err := s.uc.GetSettingByName(ctx, req.Name)
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	pbRes := &pb.Setting{
		Id:        uint32(res.ID),
		Name:      *res.Name,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}

	if res.Value != nil {
		pbRes.Value = *res.Value
	}

	return pbRes, nil
}

func (s *SettingRouter) CreateSetting(ctx context.Context, req *pb.CreateSettingReq) (*pb.Setting, error) {
	res, err := s.uc.CreateSetting(ctx, &entity.Setting{
		Name:  &req.Name,
		Value: (*datatypes.JSON)(&req.Value),
	})
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	pbRes := &pb.Setting{
		Id:        uint32(res.ID),
		Name:      *res.Name,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}

	if res.Value != nil {
		pbRes.Value = *res.Value
	}

	return pbRes, nil
}
