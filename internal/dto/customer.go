package dtos

type RegisterRequest struct {
	Name       string `json:"name" validate:"required"`
	Nik        string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type RegisterResponse struct {
	AccountNo string `json:"account_no"`
}
