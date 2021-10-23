package shop

import (
	"context"

	"github.com/TrashRaiders/garbage-truck/pkg/protos"
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	protos.UnimplementedShopServiceServer
}

func (s *Service) CreateShop(ctx context.Context, shop *protos.Shop) (*protos.Reply, error) {
	mockReply := &protos.Reply{
		Reply: &protos.Reply_Id{
			Id: "some-id",
		},
	}

	return mockReply, nil
}
func (s *Service) GetShop(ctx context.Context, req *protos.ShopRequest) (*protos.Shop, error) {
	mockShop := &protos.Shop{
		Id:          "some-id",
		OwnerId:     "some-owner",
		Name:        "some-name",
		Description: "some-description",
		Type:        "some-type",
	}

	return mockShop, nil
}
