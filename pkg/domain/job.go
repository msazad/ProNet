package domain


type Job struct{
	ID uint
	Title string
	Description string
	PostedOn string
	TotalApplication int
	CompanyName string
	PostedByID uint
}