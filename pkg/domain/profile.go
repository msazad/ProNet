package domain

type Profile struct {
	ID                uint `json:"id" gorm:"primarykey"`
	ApplicantID       uint
	ResumeFileAddress string
	Skills            string
	Education         string
	Experience        string
	Name              string
	Email             string
	Phone             string
}
