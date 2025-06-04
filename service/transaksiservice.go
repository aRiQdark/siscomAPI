package service

import (
	"errors"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"strings"

	"github.com/google/uuid"
)

type transaksiService struct {
	repository repository.Repositorytransaction
}

type Transaksiservice interface {
	// Storetransaction(transaksi models.) (models.Transactionmodel, error)
	// Findtransaction() (models.Transactionmodel, error)
	// Updatetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error)
	// Deletetransaction(transaksi models.Transactionmodel) (models.Transactionmodel, error)
	Storetransaction(input models.Transaksiinput) (models.Transactionmodel, error)
	Readdatatransaksi() ([]models.Transactionmodel, error)
}

func Newtransaction(repo repository.Repositorytransaction) Transaksiservice {
	return &transaksiService{repository: repo}
}

func (a *transaksiService) Storetransaction(input models.Transaksiinput) (models.Transactionmodel, error) {

	uid := strings.ToUpper(uuid.New().String())
	modeltransaksii := models.Transactionmodel{
		Userid:        uid,
		Nama:          input.Nama,
		TransactionID: "yes",
		Contractid:    "rets",
		Nominal:       input.Nominal,
	}
	createtransaksi, err := a.repository.Storetransaction(modeltransaksii)
	if err != nil {
		return models.Transactionmodel{}, errors.New("ada kesalahan")
	}
	return createtransaksi, err
}

func (a *transaksiService) Readdatatransaksi() ([]models.Transactionmodel, error) {
	return a.repository.Findtransaction()
}
