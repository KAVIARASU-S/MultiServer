package main

import (
	"context"
	"log"
	"net"
	"server/config"
	"server/constants"
	"server/controllers"
	"server/proto"
	"server/services"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)



func initCompany (mongoClient *mongo.Client){
	collection := config.GetCollection(mongoClient,constants.DatabaseName,"Company")
	companyService := services.Initservices(collection)
	controllers.Initcontroller(companyService)
}



func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := config.ConnectDatabase(ctx)
	defer mongoClient.Disconnect(ctx)

	if err != nil {
		log.Printf("Not Connected to Database! Resolve the issue!!!")
	}

	initCompany(mongoClient)

	lis,err := net.Listen("tcp", ":2000")
	if err !=nil {
		log.Fatalf("Failed to listen: %v", err)
	}	

	grpcServer := grpc.NewServer()

	proto.RegisterCompanyServiceServer(grpcServer,&controllers.Server{})

	log.Println("Server is running!!!")

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}