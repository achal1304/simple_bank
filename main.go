package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/achal1304/simple_bank/api"
	db "github.com/achal1304/simple_bank/db/sqlc"
	"github.com/achal1304/simple_bank/gapi"
	"github.com/achal1304/simple_bank/pb"
	"github.com/achal1304/simple_bank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations")
	}
	conn, err := sql.Open(config.DBdriver, config.DBsoruce)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	store := db.NewStore(conn)

	runGatewayServer(config, store)
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
