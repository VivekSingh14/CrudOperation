package main

import (
	"context"
	"log"
	"os"

	pb "github.com/CrudOperation/proto"

	"github.com/CrudOperation/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50052"

func main() {
	a := app.App{}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBookServiceClient(conn)

	doGet(c)

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")

}

func doGet(c pb.BookServiceClient) {
	log.Println("doGet is invoked...")

	res, err := c.Book(context.Background(), &pb.BookRequest{
		Name: "Vivek",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Name Created: %s\n", res.BookName)
}
