package repository

import (
	"errors"
	"fmt"
	models "gin-gonic-gorm/model"
	"log"

	"gorm.io/gorm"
)

type RepositoryBarang interface {
	Getallbarang() ([]models.Barang, error)
	GetbarangByID(ID string) (models.Barang, error)
	Createbarang(barang models.Barang) (models.Barang, error)
	Createcategpri(categori models.Kategori) (models.Kategori, error)
	GetcategoriByID(ID string) (models.Kategori, error)
	UpdatebarangID(barang models.Barang) (models.Barang, error)
	DeletebarangbyID(id string) error
	DeleteBarangsByIDs(ids []uint) (int64, error)
	Getallkategori() ([]models.Kategori, error)
	KelompokbykategoriID(ID string) ([]models.Kelompok, error)
	BulkDeleteByIDs(ids []string) error
}

type repositoryBarang struct {
	db *gorm.DB
}

func NewRepositoryBarang(db *gorm.DB) RepositoryBarang {
	return &repositoryBarang{db: db}
}

func (r *repositoryBarang) Getallbarang() ([]models.Barang, error) {
	var barang []models.Barang

	if err := r.db.
		Order("created_at DESC"). // Add default ordering
		Find(&barang).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Barang{}, nil // Return empty slice if no records found
		}
		return nil, fmt.Errorf("failed to get all barang: %w", err)
	}

	return barang, nil
}

func (r *repositoryBarang) GetbarangByID(ID string) (models.Barang, error) {
	var barang models.Barang
	err := r.db.Where("id = ?", ID).First(&barang).Error
	return barang, err
}

func (r *repositoryBarang) Createbarang(barang models.Barang) (models.Barang, error) {
	err := r.db.Create(&barang).Error
	if err != nil {
		log.Printf("Error create barang: %v", err)
	}
	return barang, err
}

func (r *repositoryBarang) Createcategpri(categori models.Kategori) (models.Kategori, error) {
	err := r.db.Create(&categori).Error
	if err != nil {
		log.Printf("Error create barang: %v", err)
	}
	return categori, err
}
func (r *repositoryBarang) Getallkategori() ([]models.Kategori, error) {
	var barang []models.Kategori
	err := r.db.Find(&barang).Error
	return barang, err
}
func (r *repositoryBarang) GetcategoriByID(ID string) (models.Kategori, error) {
	var Kategori models.Kategori
	err := r.db.Where("id = ?", ID).First(&Kategori).Error
	return Kategori, err
}

func (r *repositoryBarang) UpdatebarangID(barang models.Barang) (models.Barang, error) {
	err := r.db.Model(&models.Barang{}).Where("id = ?", barang.Id).Updates(barang).Error
	if err != nil {
		return barang, err
	}
	return barang, nil
}

func (a *repositoryBarang) DeletebarangbyID(id string) error {
	err := a.db.Where("id = ?", id).Delete(&models.Barang{}).Error
	return err
}

func (a *repositoryBarang) DeleteBarangsByIDs(ids []uint) (int64, error) {
	result := a.db.Where("id IN ?", ids).Delete(&models.Barang{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (r *repositoryBarang) KelompokbykategoriID(ID string) ([]models.Kelompok, error) {
	var kelompoks []models.Kelompok
	err := r.db.Where("idkategori = ?", ID).Find(&kelompoks).Error
	return kelompoks, err
}

func (r *repositoryBarang) BulkDeleteByIDs(ids []string) error {
	return r.db.Where("id IN ?", ids).Delete(&models.Barang{}).Error
}
