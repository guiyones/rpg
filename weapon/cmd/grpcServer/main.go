package main

import (
	"database/sql"
	"net"
	"weapon/internal/database"
	"weapon/internal/proto"
	"weapon/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./weapons.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	weaponDB := database.NewWeapon(db)
	weaponService := service.NewWeaponService(*weaponDB)
	grpcServer := grpc.NewServer()

	proto.RegisterWeaponServiceServer(grpcServer, weaponService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
