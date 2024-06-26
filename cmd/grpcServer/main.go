package main

import (
	"database/sql"
	"net"

	database "github.com/Mazzael/gRPC/internal"
	"github.com/Mazzael/gRPC/internal/pb"
	"github.com/Mazzael/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	courseDb := database.NewCourse(db)
	courseService := service.NewCourseService(*courseDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	pb.RegisterCourseServiceServer(grpcServer, courseService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
