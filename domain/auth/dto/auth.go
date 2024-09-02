package dto

type SignIn struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Register struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required"`
	Password     string `json:"password" validate:"required"`
}
