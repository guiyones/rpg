package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"shield/internal/database"
	"shield/internal/proto"
	"shield/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	url := "libsql://shield-guinoyo.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTIwMTc2NTQsImlkIjoiMjM1YzRlNjMtMzkyNC00NDBiLWFjNmItNWY0NjAzZjEyYTI1In0.L2gN7z5e_tJ93G8hcyADF8yXUzkeH9yXD2MSPJaF3sefTULE86SzO_F5cTk72Ew_bUDpLImTCbws1JXfOiDPDQ"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
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
