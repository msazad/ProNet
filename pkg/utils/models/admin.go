package models

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=6,max=12"`
}
type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type AdminToken struct {
	Admin        AdminDetailsResponse
	Token        string
	RefreshToken string
}