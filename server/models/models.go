package models

type Company struct{
	CompanyName string `json:"CompanyName" bson:"CompanyName"`
	CEO string `json:"CEO" bson:"CEO"`
}