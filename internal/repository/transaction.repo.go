package repository

import (
	"github.com/savanyv/account-service-api/internal/models"
	"github.com/savanyv/account-service-api/internal/utils"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByAccountNo(accountNo string) ([]*models.Transaction, error)
	FindByAccountNoAndType(accountNo, transactionType string) ([]*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		utils.LogError("REPOSITORY", "Failed to create transaction: %v", err)
		return err
	}
	return nil
}

func (r *transactionRepository) FindByAccountNo(accountNo string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	if err := r.db.Where("account_no = ?", accountNo).Order("created_at desc").Find(&transactions).Error; err != nil {
		utils.LogError("REPOSITORY", "Failed to find transactions: %v", err)
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) FindByAccountNoAndType(accountNo, transactionType string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	if err := r.db.Where("account_no = ? AND type = ?", accountNo, transactionType).Order("created_at DESC").Find(&transactions).Error; err != nil {
		utils.LogError("REPOSITORY", "Failed to find transactions: %v", err)
		return nil, err
	}

	return transactions, nil
}
