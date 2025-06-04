package service

import (
	"gin-gonic-gorm/entity"
	// models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
)

type Accountservice interface {
	Findallaccount() ([]entity.Account, error)
	FindaccountByID(accountID string) ([]entity.Account, error)
	FindaccountBymemberID(memberID string) ([]entity.Account, error)
}

type accountService struct {
	repository repository.AccountRepo
}

func NewServiceaccount(repo repository.AccountRepo) Accountservice {
	return &accountService{repository: repo}
}


func (s *accountService) FindaccountByID(accountID string) ([]entity.Account, error) {
	return s.repository.Findallaccount()
}

// FindaccountBymemberID implements Accountservice
func (s *accountService) FindaccountBymemberID(memberID string) ([]entity.Account, error) {
	panic("unimplemented")
}

// Findallaccount implements Accountservice
func (s *accountService) Findallaccount() ([]entity.Account, error) {
	panic("unimplemented")
}

