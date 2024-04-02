package service

import (
	"context"
	"shield/internal/database"
	"shield/internal/proto"
)

type ShieldService struct {
	proto.UnimplementedShieldServiceServer
	ShieldDB database.Shield
}

func NewShieldService(db database.Shield) *ShieldService {
	return &ShieldService{
		ShieldDB: db,
	}
}

func (s *ShieldService) CreateShield(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {

	shield, err := s.ShieldDB.CreateShield(in.Name, in.Defense)
	if err != nil {
		return nil, err
	}

	shieldResponse := &proto.Shield{
		Name:    shield.Name,
		Defense: shield.Defense,
	}

	return &proto.CreateResponse{
		Shield: shieldResponse,
	}, nil
}

func (s *ShieldService) ListShields(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {

	shields, err := s.ShieldDB.GetAllShields()
	if err != nil {
		return nil, err
	}

	var shieldsResponse []*proto.Shield

	for _, shield := range shields {
		shieldResponse := &proto.Shield{
			Id:      shield.ID,
			Name:    shield.Name,
			Defense: shield.Defense,
		}

		shieldsResponse = append(shieldsResponse, shieldResponse)
	}

	return &proto.GetListResponse{
		Shields: shieldsResponse,
	}, nil
}
