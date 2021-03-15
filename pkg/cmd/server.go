package cmd

import (
	"fmt"
	protos "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/grpc/note"
	note "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/module/user"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/storage/postgress"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/service"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	postgresURL = "postgres://%s:%s@%s/%s?sslmode=disable"

	dialect = "postgres"
)

func RunServer() error {

	dbURL := fmt.Sprintf(postgresURL, "postgres", "root", "localhost", "Note")

	dbConn, err := gorm.Open(dialect, dbURL)
	if dbConn != nil {
		defer dbConn.Close()
	}
	if err != nil {
		panic(err)
		return err
	}
	//createTable(dbConn)
	postgresUser := postgress.NewNoteRepository(dbConn)
	srv := note.NewService(*postgresUser)

	gs := grpc.NewServer()

	c := service.NewNoteServer(srv)

	protos.RegisterNoteServiceServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		return err
	}

	err = gs.Serve(l)
	if err != nil {
		return err
	}
	return nil
}
