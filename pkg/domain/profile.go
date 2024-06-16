package domain

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID            uint   `json:"user"`
	ResumeFileAddress string `json:"resume_file_address" gorm:"not null"`
	Skills            string `json:"skills" gorm:"not null"`
	Education         string `json:"education" gorm:"not null"`
	Experience        string `json:"experience" gorm:"not null"`
	Name              string `json:"name" gorm:"validation required"`
	Email             string `json:"email" gorm:"not null"`
	Phone             string `json:"phone" gorm:"not null"`
}
