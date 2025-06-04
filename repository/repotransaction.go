package repository

import (
	models "gin-gonic-gorm/model"

	"gorm.io/gorm"
)

type Repositorytransaction interface {
	Storetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error)
	Findtransaction() ([]models.Transactionmodel, error)
	Updatetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error)
	Deletetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error)
}

type repositoryT struct {
	db *gorm.DB
}

func NewRepository1(db *gorm.DB) Repositorytransaction {
	return &repositoryT{db: db}
}

func (a *repositoryT) Storetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error) {
	err := a.db.Create(&transaksi)
	if err != nil {
		return transaksi, nil
	}
	return transaksi, nil
}

func (a *repositoryT) Findtransaction() ([]models.Transactionmodel, error) {
	var transaksi []models.Transactionmodel
	err := a.db.Find(&transaksi).Error
	return transaksi, err
}

func (a *repositoryT) Updatetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error) {
	err := a.db.Model(&models.Transactionmodel{}).Where("TransactionID = ?", transaksi.TransactionID).Updates(transaksi).Error
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}

func (a *repositoryT) Deletetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error) {
	err := a.db.Model(&models.Transactionmodel{}).Where("TransactionID = ?", transaksi.TransactionID).Delete(transaksi).Error
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}


