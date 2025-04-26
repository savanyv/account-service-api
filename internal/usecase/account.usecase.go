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
	Deposit(req *dtos.DepositRequest) (*dtos.DepositResponse, error)
	Withdraw(req *dtos.WithdrawRequest) (*dtos.WithdrawResponse, error)
	GetBalance(accountNo string) (*dtos.BalanceResponse, error)
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

func (u *accountUsecase) Deposit(req *dtos.DepositRequest) (*dtos.DepositResponse, error) {
	// find customer
	customer, err := u.accountRepo.FindByAccountNo(req.AccountNo)
	if err != nil {
		utils.LogError("USECASE", "Failed to find customer: %v", err)
		return nil, err
	}

	// create transaction
	transaction := &models.Transaction{
		AccountNo: req.AccountNo,
		Type: "deposit",
		Amount: req.Amount,
		FinalBalance: customer.Balance + req.Amount,
		CreatedAt: time.Now(),
	}

	// save transaction
	if err := u.transactionRepo.Create(transaction); err != nil {
		utils.LogError("USECASE", "Failed to create transaction: %v", err)
		return nil, err
	}

	// update customer balance
	customer.Balance += req.Amount
	if err := u.accountRepo.Update(customer); err != nil {
		utils.LogError("USECASE", "Failed to update customer: %v", err)
		return nil, err
	}

	// response
	resp := &dtos.DepositResponse{
		AccountNo: customer.AccountNo,
		Balance: customer.Balance,
	}

	return resp, nil
}

func (u *accountUsecase) Withdraw(req *dtos.WithdrawRequest) (*dtos.WithdrawResponse, error) {
	// find customer
	customer, err := u.accountRepo.FindByAccountNo(req.AccountNo)
	if err != nil {
		utils.LogError("USECASE", "Failed to find customer: %v", err)
		return nil, err
	}

	// create transaction
	transaction := &models.Transaction{
		AccountNo: req.AccountNo,
		Type: "withdraw",
		Amount: req.Amount,
		FinalBalance: customer.Balance - req.Amount,
		CreatedAt: time.Now(),
	}

	// save transaction
	if err := u.transactionRepo.Create(transaction); err != nil {
		utils.LogError("USECASE", "Failed to create transaction: %v", err)
		return nil, err
	}

	// update customer balance
	customer.Balance -= req.Amount
	if err := u.accountRepo.Update(customer); err != nil {
		utils.LogError("USECASE", "Failed to update customer: %v", err)
		return nil, err
	}

	// response
	resp := &dtos.WithdrawResponse{
		AccountNo: customer.AccountNo,
		Balance: customer.Balance,
	}

	return resp, nil
}

func (u *accountUsecase) GetBalance(accountNo string) (*dtos.BalanceResponse, error) {
	// find customer
	customer, err := u.accountRepo.FindByAccountNo(accountNo)
	if err != nil {
		utils.LogError("USECASE", "Failed to find customer: %v", err)
		return nil, err
	}

	// response
	resp := &dtos.BalanceResponse{
		Balance: customer.Balance,
	}

	return resp, nil
}
