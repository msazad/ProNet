package domain

import (
	"fmt"
	"linkdin/pkg/domain"
	"linkdin/pkg/repository/interfaces"
	"linkdin/pkg/utils/models"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func newAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ar *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
	var adminCompareDetails domain.Admin
	err := ar.DB.Raw("select * from admins where user_name=?", adminDetails.Email).Scan(&adminCompareDetails).Error
	if err != nil {
		return domain.Admin{}, err
	}
	fmt.Println("admin compare details from repo", adminCompareDetails)
	return adminCompareDetails,nil
}
