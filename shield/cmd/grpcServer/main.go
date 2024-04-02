package main

import (
	"database/sql"
	"net"
	"shield/internal/database"
	"shield/internal/proto"
	"shield/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./shield.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	shieldDB := database.NewShield(db)
	shieldService := service.NewShieldService(*shieldDB)
	grpcServer := grpc.NewServer()

	proto.RegisterShieldServiceServer(grpcServer, shieldService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
