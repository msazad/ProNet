package interfaces

import (
	"linkdin/pkg/domain"
	"linkdin/pkg/utils/models"
)


type AdminRepository interface{
	LoginHandler(adminDetails models.AdminLogin)(domain.Admin,error)
}