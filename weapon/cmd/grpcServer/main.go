package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"weapon/internal/database"
	"weapon/internal/proto"
	"weapon/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	url := "libsql://weapon-guinoyo.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTIwMTc2OTcsImlkIjoiOTY2YmEwZmQtYTA1YS00OWEyLWFmNzAtYzkyZTIwMjRmODVlIn0.wbOeSc0WLdkJECixRax8J62vOqwnNRWi2ek1eHcOuWoHLrkG4-Tz5FT2XjKgKN_sIbePmY1_tXBJ3BASGqeFBw"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
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
