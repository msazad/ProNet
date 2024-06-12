package domain

type User struct {
	ID             uint
	Name           string
	Email          string
	Address        string
	UserType       int
	PasswordHash   string
	ProfileHeading string
	Profile        Profile
}
