package service

import (
	"context"
	"weapon/internal/database"
	"weapon/internal/proto"
)

type WeaponService struct {
	proto.UnimplementedWeaponServiceServer
	WeaponDB database.Weapon
}

func NewWeaponService(weaponDB database.Weapon) *WeaponService {
	return &WeaponService{
		WeaponDB: weaponDB,
	}
}

func (s *WeaponService) CreateWeapon(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {

	weapon, err := s.WeaponDB.CreateWeapon(in.Name, in.Distance, in.Category, in.WeaponRange, in.Hand, in.Damage, in.Defense)
	if err != nil {
		return nil, err
	}

	weaponResponse := &proto.Weapon{
		Name:        weapon.Name,
		Distance:    weapon.Distance,
		Category:    weapon.Category,
		WeaponRange: weapon.WeaponRange,
		Hand:        weapon.Hand,
		Damage:      weapon.Damage,
		Defense:     weapon.Defense,
	}

	return &proto.CreateResponse{
		Weapon: weaponResponse,
	}, nil
}

func (s *WeaponService) ListWeapons(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {

	weapons, err := s.WeaponDB.GetAllWeapon()
	if err != nil {
		return nil, err
	}

	var weaponsResponse []*proto.Weapon

	for _, weapon := range weapons {
		weaponResponse := &proto.Weapon{
			Id:          weapon.ID,
			Name:        weapon.Name,
			Distance:    weapon.Distance,
			Category:    weapon.Category,
			WeaponRange: weapon.WeaponRange,
			Hand:        weapon.Hand,
			Damage:      weapon.Damage,
			Defense:     weapon.Defense,
		}

		weaponsResponse = append(weaponsResponse, weaponResponse)
	}

	return &proto.GetListResponse{
		Weapons: weaponsResponse,
	}, nil

}

func (s *WeaponService) GetByCategory(ctx context.Context, in *proto.GetByCategoryRequest) (*proto.GetListResponse, error) {

	weapons, err := s.WeaponDB.GetByCategory(in.Category)
	if err != nil {
		return nil, err
	}

	var weaponsResponse []*proto.Weapon

	for _, weapon := range weapons {
		weaponResponse := &proto.Weapon{
			Id:          weapon.ID,
			Name:        weapon.Name,
			Distance:    weapon.Distance,
			Category:    weapon.Category,
			WeaponRange: weapon.WeaponRange,
			Hand:        weapon.Hand,
			Damage:      weapon.Damage,
			Defense:     weapon.Defense,
		}

		weaponsResponse = append(weaponsResponse, weaponResponse)
	}

	return &proto.GetListResponse{
		Weapons: weaponsResponse,
	}, nil
}
