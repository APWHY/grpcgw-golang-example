package services

import (
	"context"
	"database/sql"
	"github.com/APWHY/grpcgw-golang-example/pet"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/APWHY/grpcgw-golang-example/proto"
)

type petService struct {
	repo pet.Repository
}

func (s *petService) GetPets(ctx context.Context, _ *empty.Empty) (*proto.GetPetsResponse, error) {
	res := &proto.GetPetsResponse{Meta: &proto.ResponsePayload{Success: true}}
	pets, err := s.repo.GetPets(ctx)

	if err != nil {
		res.Meta.Error = err.Error()
		res.Meta.Success = false
		return res, err
	}

	for _, pet := range pets {
		res.Pets = append(res.Pets, repoToProto(pet))
	}
	return res, nil
}

func (s *petService) PostPet(ctx context.Context, req *proto.PostPetRequest) (*proto.PostPetResponse, error) {
	res := &proto.PostPetResponse{Meta: &proto.ResponsePayload{Success: true}}
	pet, err := s.repo.InsertPet(ctx, protoToRepo(req.Pet))

	if err != nil {
		res.Meta.Error = err.Error()
		res.Meta.Success = false
		return nil, err
	}
	res.Pet = repoToProto(pet)
	return res, nil
}

// NewPetService - Returns new implementation of pet service
func NewPetService(repository pet.Repository) proto.PetServiceServer {
	return &petService{repo: repository}
}

func repoToProto(in pet.Pet) *proto.Pet {
	return &proto.Pet{
		Id:      in.ID,
		Name:    in.Name,
		Owner:   in.Owner,
		Species: in.Species,
		Sex:     in.Sex,
		Birth:   in.Birth,
		Death:   in.Death.String,
		Created: in.Created.String(),
		Updated: in.Updated.String(),
	}
}

func protoToRepo(in *proto.Pet) pet.Pet {
	return pet.Pet{
		ID:      in.Id,
		Name:    in.Name,
		Owner:   in.Owner,
		Species: in.Species,
		Sex:     in.Sex,
		Birth:   in.Birth,
		Death:   sql.NullString{String: in.Death, Valid: bool(len(in.Death) != 0)},
	}
}
