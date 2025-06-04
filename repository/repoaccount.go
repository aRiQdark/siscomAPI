package repository

import (
	"gin-gonic-gorm/entity"

	"gorm.io/gorm"
)

type AccountRepo interface {
	Findallaccount() ([]entity.Account, error)
	FindaccountByID(accountID string) ([]entity.Account, error)
	FindacoountBymemberID(memberID string) ([]entity.Account,error)
}

type accounts struct {
	db *gorm.DB
}

func NewRepositoryaccount(db *gorm.DB) Repository {
	return &repository{db: db}
}



func (r *accounts) FindaccountByID(accountID string) ([]entity.Account, error) {
	var account []entity.Account
	err := r.db.Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, err
}

func (r *accounts) Findallaccount(accountID string) ([]entity.Account, error) {
	var account []entity.Account
	err := r.db.Where("accoung_id = ?",account).First(&account).Error
	return account,err
}

func (r *accounts)FindacountBymemberID(memberID string) ([]entity.Account,error) {
	var account []entity.Account
	err := r.db.Where("member_id = ?",memberID).First(&account).Error
	return account,err
}
