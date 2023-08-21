package gapi

import (
	"testmongo/dbmongo"
	"testmongo/gapi/pb"
)

type Server struct {
	pb.UnimplementedTestGrpcServer
	mongoStore dbmongo.MongoStore
}

func NewServer(mongoStore dbmongo.MongoStore) *Server {
	server := &Server{
		mongoStore: mongoStore,
	}

	return server
}
