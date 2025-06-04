package repository

import (
	"errors"
	models "gin-gonic-gorm/model"

	"gorm.io/gorm"
)

type Pesananrepository interface {
	Createpesanan(models models.Modelpesanan) (models.Modelpesanan, error)
	Findallpesanan() ([]models.Modelpesanan, error)
	FindpesananBymemberid(memberID string) ([]models.Modelpesanan, error)
}

type repositorypesanan struct {
	db *gorm.DB
}

func NewrepositoryPesanan(db *gorm.DB) Pesananrepository {
	return &repositorypesanan{db: db}
}

func (r *repositorypesanan) Createpesanan(models models.Modelpesanan) (models.Modelpesanan, error) {
	err := r.db.Create(&models).Error
	if err != nil {
		return models, errors.New("Gagal membuat pesanan,silahkan coba kembali")
	}
	return models, err
}

func (r *repositorypesanan) Findallpesanan() ([]models.Modelpesanan, error) {
	var models []models.Modelpesanan

	err := r.db.Find(&models).Error
	return models, err
}

func (r *repositorypesanan) FindpesananBymemberid(memberID string) ([]models.Modelpesanan, error) {
	var pesanan []models.Modelpesanan
	err := r.db.Where("member_id = ?", memberID).Find(&pesanan).Error

	return pesanan, err
}
