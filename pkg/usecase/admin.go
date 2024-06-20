package usecase

import (
	"fmt"
	"linkdin/pkg/domain"
	"linkdin/pkg/helper"
	"linkdin/pkg/repository/interfaces"
	services "linkdin/pkg/usecase/interfaces"
	"linkdin/pkg/utils/models"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUsecase(adRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUsecase{
		adminRepository: adRepo,
	}
}

func (au *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error) {
	//Getting details of the admin based on the email provided
	adminCompareDetails, _ := au.adminRepository.LoginHandler(adminDetails)

	if adminCompareDetails.Email != adminDetails.Email {
		return domain.AdminToken{}, errors.New("admin not exist")

	}
	//Compare password from database that provided by admin
	err := bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.PasswordHash), []byte(adminDetails.Password))
	if err != nil {
		return domain.AdminToken{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse
	adminDetailsResponse.ID = int(adminCompareDetails.ID)
	adminDetailsResponse.Name = adminCompareDetails.Name
	adminDetailsResponse.Email = adminCompareDetails.Email
	fmt.Println("reached the admindetailsresponse")
	token, refresh, err := helper.GenerateAdminToken(adminDetailsResponse)
	if err != nil {
		return domain.AdminToken{}, err
	}
	return domain.AdminToken{
		Admin:        adminDetailsResponse,
		Token:        token,
		RefreshToken: refresh,
	}, nil
}
