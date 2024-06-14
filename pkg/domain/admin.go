package domain

import (
	"linkdin/pkg/utils/models"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name         string  `json:"name" gorm:"validate:required"`
	Email        string  `json:"email" gorm:"unique"`
	UserType     bool    `json:"user_type"`
	PasswordHash string  `json:"password_hash"`
	Profile      Profile `json:"profile"`
	JobsPosted   []Job   `json:"jobs_posted"`
	Users        []User  `json:"user"`
}
type AdminToken struct{
	Admin models.AdminDetailsResponse
	Token string
	RefreshToken string
}
