package controllers

import (
	"context"
	"log"
	"server/interfaces"
	"server/models"
	"server/proto"
)

var	CompanyService interfaces.Icompany


type Server struct {
	proto.UnimplementedCompanyServiceServer
}


func Initcontroller (companyService interfaces.Icompany) {
	CompanyService =companyService
}



func (controller *Server) DisplayCompany (ctx context.Context, empty *proto.Empty)(*proto.CompanyResponse,error){
	allCompany,err := CompanyService.DisplayCompany()

	if err != nil {
		return nil,err
	}

	var protoarray []*proto.Company

	for _,v := range *allCompany{
		var protocompany proto.Company
		protocompany.CompanyName = v.CompanyName
		protocompany.CEO = v.CEO
		protoarray=append(protoarray, &proto.Company{
			CompanyName: protocompany.CompanyName,
			CEO:         protocompany.CEO,
		})
	}

	response := &proto.CompanyResponse{
		Companies: protoarray,
	}

	return response,nil
}

func (controller *Server) InsertCompany (ctx context.Context, protoCompany *proto.Company) (*proto.Empty,error){
	log.Println("Insert Company started")
	
	
	var company models.Company
    //var empty *proto.Empty
	company.CompanyName = protoCompany.CompanyName
	log.Println("error")
	company.CEO = protoCompany.CEO
    
	log.Println("Going to service")
	if err := CompanyService.InsertCompany(&company); err != nil {
		return nil,err
	}
    empty:= &proto.Empty{}
	return empty,nil

}