package controllers

import (
	"client/models"
	"client/proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


var Client proto.CompanyServiceClient


func Initcontroller (client proto.CompanyServiceClient) {
	Client = client
}

func DisplayCompany (c *gin.Context){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	req:= &proto.Empty{}

	allCompany,err := Client.DisplayCompany(ctx,req)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error in displaying Company Names":err,
		})
		return
	}
	
	c.JSON(http.StatusOK,allCompany)
}

func InsertCompany (c *gin.Context){
	var company *models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	req := &proto.Company{
		CompanyName: company.CompanyName,
		CEO:         company.CEO,
	}
	insertCompany,err := Client.InsertCompany(ctx,req)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error in grpc server":err,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Success",
		"Inserted":insertCompany,
	})

}