package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mauricio-pagarme/go-grpc-example/internal/database"
	"github.com/mauricio-pagarme/go-grpc-example/internal/pb"
	"github.com/mauricio-pagarme/go-grpc-example/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userDb := database.NewUser(db)
	transactionDb := database.NewTransactionType(db)

	userService := service.NewUserService(*userDb)                              // configuro o banco de dados no meu service
	transactionTypeService := service.NewTransactionTypeService(*transactionDb) // configuro o banco de dados no meu service

	grpcServer := grpc.NewServer()                                              // subo um servidor grpc
	pb.RegisterUserServiceServer(grpcServer, userService)                       // registro meu serviço no grpc
	pb.RegisterTransactionTypeServiceServer(grpcServer, transactionTypeService) // registro meu serviço no grpc

	reflection.Register(grpcServer) // para teste vamos usar reflection com o evans

	lis, err := net.Listen("tcp", ":50051") // habilito um listener tcp
	if err != nil {
		panic(err)
	}
	println("server running")
	if err := grpcServer.Serve(lis); err != nil { // cadastro esse tcp no meu grpc serve
		panic(err)
	}

}
