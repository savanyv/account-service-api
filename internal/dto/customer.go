package dtos

type RegisterRequest struct {
	Name       string `json:"name" validate:"required"`
	Nik        string `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type RegisterResponse struct {
	AccountNo string `json:"account_no"`
}

type DepositRequest struct {
	AccountNo string `json:"account_no" validate:"required"`
	Amount    int64  `json:"amount" validate:"required"`
}

type DepositResponse struct {
	AccountNo string `json:"account_no"`
	Balance int64 `json:"balance"`
}

type WithdrawRequest struct {
	AccountNo string `json:"account_no" validate:"required"`
	Amount int64 `json:"amount" validate:"required"`
}

type WithdrawResponse struct {
	AccountNo string `json:"account_no"`
	Balance int64 `json:"balance"`
}
