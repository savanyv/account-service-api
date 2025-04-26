package repository

import (
	"errors"

	"github.com/savanyv/account-service-api/internal/models"
	"github.com/savanyv/account-service-api/internal/utils"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(customer *models.Customer) error
	FindByNIK(nik string) (*models.Customer, error)
	FindByPhone(phone string) (*models.Customer, error)
	FindByAccountNo(accountNo string) (*models.Customer, error)
	Update(customer *models.Customer) error
	Begin() AccountRepository
	Commit() error
	Rollback()
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) Create(customer *models.Customer) error {
	// Start transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if NIK exists
	if existing, err := r.FindByNIK(customer.Nik); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	} else if existing != nil {
		tx.Rollback()
		return errors.New("NIK already registered")
	}

	// Check if phone exists
	if existing, err := r.FindByPhone(customer.PhoneNumber); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	} else if existing != nil {
		tx.Rollback()
		return errors.New("phone number already registered")
	}

	// Create customer
	if err := tx.Create(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *accountRepository) FindByNIK(nik string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("nik = ?", nik).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &customer, nil
}

func (r *accountRepository) FindByPhone(phone string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("phone_number = ?", phone).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &customer, nil
}

func (r *accountRepository) FindByAccountNo(accountNo string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("account_no = ?", accountNo).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &customer, nil
}

func (r *accountRepository) Update(customer *models.Customer) error {
	err := r.db.Save(customer).Error
	if err != nil {
		utils.LogError("REPOSITORY", "Failed to update customer: %v", err)
		return err
	}
	return nil
}

func (r *accountRepository) Begin() AccountRepository {
    return &accountRepository{
	  db: r.db.Begin(),
    }
}

func (r *accountRepository) Commit() error {
    return r.db.Commit().Error
}

func (r *accountRepository) Rollback() {
    r.db.Rollback()
}
