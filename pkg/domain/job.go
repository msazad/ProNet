package domain

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	PostedOn         time.Time `json:"posted_on"`
	TotalApplication int       `json:"total_application"`
	CompanyName      string    `json:"company_name"`
	PostedByID       uint      `json:"posted_by_id"`
	PostedBy         User      `json:"posted_by"`
}
