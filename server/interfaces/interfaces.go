package interfaces

import "server/models"



type Icompany interface {
	DisplayCompany () (*[]models.Company,error)
	InsertCompany (company *models.Company) (err error)
}