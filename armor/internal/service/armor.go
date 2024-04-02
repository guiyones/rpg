package service

import (
	"armor/internal/database"
	"armor/internal/proto"
	"context"
)

type ArmorService struct {
	proto.UnimplementedArmorServiceServer
	ArmorDB database.Armor
}

func NewArmorService(armorDB database.Armor) *ArmorService {
	return &ArmorService{
		ArmorDB: armorDB,
	}
}

func (s *ArmorService) CreateArmor(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {

	armor, err := s.ArmorDB.CreateArmor(in.Name, in.Category, in.Defense)
	if err != nil {
		return nil, err
	}

	armorResponse := &proto.Armor{
		Name:     armor.Name,
		Category: armor.Category,
		Defense:  armor.Defense,
	}

	return &proto.CreateResponse{
		Armor: armorResponse,
	}, nil
}

func (s *ArmorService) ListArmors(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {

	armors, err := s.ArmorDB.GetAllArmors()
	if err != nil {
		return nil, err
	}

	var armorsResponse []*proto.Armor

	for _, armor := range armors {
		armorResponse := &proto.Armor{
			Id:       armor.ID,
			Name:     armor.Name,
			Category: armor.Category,
			Defense:  armor.Defense,
		}

		armorsResponse = append(armorsResponse, armorResponse)
	}

	return &proto.GetListResponse{
		Armors: armorsResponse,
	}, nil
}

func (s *ArmorService) GetByCategory(ctx context.Context, in *proto.GetByCategoryRequest) (*proto.GetListResponse, error) {

	armors, err := s.ArmorDB.GetByCategory(in.Category)
	if err != nil {
		return nil, err
	}

	var armorsResponse []*proto.Armor

	for _, armor := range armors {
		armorResponse := &proto.Armor{
			Id:       armor.ID,
			Name:     armor.Name,
			Category: armor.Category,
			Defense:  armor.Defense,
		}

		armorsResponse = append(armorsResponse, armorResponse)
	}

	return &proto.GetListResponse{
		Armors: armorsResponse,
	}, nil
}
