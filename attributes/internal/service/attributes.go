package service

import (
	"attributes/internal/database"
	"attributes/internal/proto"
	"context"
)

type AttributesService struct {
	proto.UnimplementedAttributesServiceServer
	AttributesDB database.Attributes
}

func NewAttributesService(attributeDB database.Attributes) *AttributesService {
	return &AttributesService{
		AttributesDB: attributeDB,
	}
}

func (s *AttributesService) CreateAttibutes(ctx context.Context, in *proto.CreateAttributeRequest) (*proto.Attribute, error) {

	attribute, err := s.AttributesDB.CreateAttibutes(in.Name, in.AttributeType, in.Description, in.Value)
	if err != nil {
		return nil, err
	}

	attributeResponse := &proto.Attribute{
		Name:          attribute.Name,
		AttributeType: attribute.AttributeType,
		Description:   attribute.Description,
		Value:         attribute.Value,
	}

	return attributeResponse, nil
}

func (s *AttributesService) ListAttibutes(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {

	attributes, err := s.AttributesDB.GetAllAtributes()
	if err != nil {
		return nil, err
	}

	var attributesResponse []*proto.Attribute

	for _, attribute := range attributes {
		attributeResponse := &proto.Attribute{
			Id:            attribute.ID,
			Name:          attribute.Name,
			AttributeType: attribute.AttributeType,
			Description:   attribute.Description,
			Value:         attribute.Value,
		}

		attributesResponse = append(attributesResponse, attributeResponse)
	}

	return &proto.GetListResponse{
		Attributes: attributesResponse,
	}, nil
}

func (s *AttributesService) GetByAttributeType(ctx context.Context, in *proto.GetByAttributeTypeRequest) (*proto.GetListResponse, error) {

	attributes, err := s.AttributesDB.GetByAttributeType(in.AttributeType)
	if err != nil {
		return nil, err
	}

	var attributesResponse []*proto.Attribute

	for _, attribute := range attributes {
		attributeResponse := &proto.Attribute{
			Id:            attribute.ID,
			Name:          attribute.Name,
			AttributeType: attribute.AttributeType,
			Description:   attribute.Description,
			Value:         attribute.Value,
		}

		attributesResponse = append(attributesResponse, attributeResponse)
	}

	return &proto.GetListResponse{
		Attributes: attributesResponse,
	}, nil
}
