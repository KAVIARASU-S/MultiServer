package main

import (
	"client/constants"
	"client/controllers"
	"client/proto"
	"client/routes"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)




func main() {
	conn , err := grpc.Dial(":2000",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

    client := proto.NewCompanyServiceClient(conn)


	controllers.Initcontroller(client)

	clientServer := gin.Default()

	routes.Routes(clientServer)	

	clientServer.Run(constants.Port)
}