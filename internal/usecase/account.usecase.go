package usecase

import (
	"errors"
	"time"

	dtos "github.com/savanyv/account-service-api/internal/dto"
	"github.com/savanyv/account-service-api/internal/models"
	"github.com/savanyv/account-service-api/internal/repository"
	"github.com/savanyv/account-service-api/internal/utils"
)

type AccountUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.RegisterResponse, error)
}

type accountUsecase struct {
	accountRepo repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func NewAccountUsecase(accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *accountUsecase) Register(req *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	// check if NIK exists
	if existing, _ := u.accountRepo.FindByNIK(req.Nik); existing != nil {
		utils.LogError("USECASE", "NIK already registered")
		return nil, errors.New("NIK already registered")
	}

	// check if phone number exists
	if existing, _ := u.accountRepo.FindByPhone(req.PhoneNumber); existing != nil {
		utils.LogError("USECASE", "Phone number already registered")
		return nil, errors.New("phone number already registered")
	}

	// create customer
	accountNo := utils.GenerateAccountNo()
	customer := &models.Customer{
		Nik: req.Nik,
		Name: req.Name,
		PhoneNumber: req.PhoneNumber,
		AccountNo: accountNo,
		Balance: 0,
		CreatedAt: time.Now(),
	}

	// save customer
	if err := u.accountRepo.Create(customer); err != nil {
		utils.LogError("USECASE", "Failed to create customer: %v", err)
		return nil, err
	}

	// response
	resp := &dtos.RegisterResponse{
		AccountNo: accountNo,
	}

	return resp, nil
}
