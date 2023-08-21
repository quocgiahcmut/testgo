package gapi

import (
	"context"

	"testmongo/dbmongo"
	"testmongo/gapi/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePerson(ctx context.Context, req *pb.CreatePersonRequest) (*pb.CreatePersonResponse, error) {
	args := &dbmongo.CreatePersonParams{
		Name: req.Name,
		Age:  int(req.Age),
		Bio:  req.Bio,
	}

	person, err := server.mongoStore.InsertPerson(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	res := &pb.CreatePersonResponse{
		Id:   person.ID.Hex(),
		Name: person.Name,
		Age:  int64(person.Age),
		Bio:  person.Bio,
	}

	return res, nil
}

func (server *Server) GetListPerson(ctx context.Context, req *pb.GetListPersonReq) (*pb.GetListPersonRes, error) {
	people, err := server.mongoStore.GetListPerson(ctx, req.Page, req.PerPage)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	newPeople := convertToPersonResponse(people)

	return &pb.GetListPersonRes{People: newPeople}, nil
}

func convertToPersonResponse(mongoPerson []*dbmongo.Person) []*pb.Person {
	var pbPerson []*pb.Person

	for _, person := range mongoPerson {
		newPerson := &pb.Person{
			Id:   person.ID.Hex(),
			Name: person.Name,
			Age:  int64(person.Age),
			Bio:  person.Bio,
		}
		pbPerson = append(pbPerson, newPerson)
	}

	return pbPerson
}
