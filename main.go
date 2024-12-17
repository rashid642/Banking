package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	db "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/api"
	"github.com/rashid642/banking/gapi"
	"github.com/rashid642/banking/pb"
	"github.com/rashid642/banking/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig(".") 
	if err != nil {
		log.Fatal("Can not load config files, err:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database, error :", err)
	}
	store := db.NewStore(conn) 
	// runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store) 
	if err != nil {
		log.Fatal("Can not create the server, err :", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankingServer(grpcServer, server)
	reflection.Register(grpcServer) // help client to get to know about function and how to call them 

	listner, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("Can not create the listner, err :", err)
	}

	log.Printf("start the gRPC server at %s", listner.Addr().String())
	err = grpcServer.Serve(listner) 
	if err != nil {
		log.Fatal("Can not start the listner, err :", err)
	}
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store) 
	if err != nil {
		log.Fatal("Can not create the server, err :", err)
	}

	err = server.Start(config.HTTPServerAddress) 
	if err != nil {
		log.Fatal("Can not start the server, err :", err)
	}
}