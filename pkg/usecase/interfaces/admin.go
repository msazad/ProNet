package interfaces

import "linkdin/pkg/utils/models"

type AdminUsecase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error)
}
