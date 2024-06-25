package handlers

import (
	"linkdin/pkg/response"
	"linkdin/pkg/utils/models"
	"net/http"

	services "linkdin/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUsecase services.AdminUsecase
}

func newAdminHandler(adminUsecase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}
func (ah *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ah.adminUsecase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "can't autheticate admin", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetCookie("Refresh", admin.RefreshToken, 3600*24*30, "/", "", false, true)
	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

