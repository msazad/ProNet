package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"validate:required"`
	Email          string `json:"email" gorm:"unique"`
	Address        string `json:"address"`
	UserType       bool `json:"user_type"`
	PasswordHash   string `json:"password_hash"`
	ProfileHeading string `json:"profile_heading"`
	Profile        Profile `json:"profile"`
	JobsPosted []Job `json:"jobs_posted"`
}
