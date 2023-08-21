package main

import (
	"context"
	"log"
	"net"

	"testmongo/api"
	"testmongo/dbmongo"
	"testmongo/gapi"
	pb "testmongo/gapi/pb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")) // context.TODO() use to create blank context
	if err != nil {
		log.Fatal("can't connect to mongo database: ", err)
	}

	mongoStore := dbmongo.NewMongoStore(client.Database("testdb"))
	RunGrpcServer(*mongoStore)
}

func RunGrpcServer(mongoStore dbmongo.MongoStore) {
	server := gapi.NewServer(mongoStore)

	grpcServer := grpc.NewServer()
	pb.RegisterTestGrpcServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("can't create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can't start gRPC Server")
	}
}

func RunApiServer(mongoStore dbmongo.MongoStore) {
	server := api.NewServer(mongoStore)

	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("can't start REST API server: ", err)
	}
}
