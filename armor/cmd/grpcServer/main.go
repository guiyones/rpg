package main

import (
	"armor/internal/database"
	"armor/internal/proto"
	"armor/internal/service"
	"database/sql"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	// db, err := sql.Open("sqlite3", "./armor.sqlite")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	url := "libsql://armor-guinoyo.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTIwMDcwNTksImlkIjoiNDJjMTQ1ZDUtMjIyZi00ZGIyLWI1MWUtNzI5NTkyMmZmY2QwIn0.hGjXVCYHgUs9egZmM79P9rxmBG6wK1RkBog79X4_Or6IVKQ9vpaC3h3beDA3Hh8EaclMBN40d2_Qe9xeLRtVAg"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	armorDB := database.NewArmor(db)
	armorService := service.NewArmorService(*armorDB)
	grpcServer := grpc.NewServer()

	proto.RegisterArmorServiceServer(grpcServer, armorService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
