package service

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"log"
	"strings"

	"github.com/google/uuid"
)

type Barangservice interface {
	Getallbarang() ([]models.Barang, error)
	Getbarangbyid(ID string) (models.Barang, error)
	Createbarang(input models.Baranginput) (models.Barang, error)
	Getallkategori() ([]models.Kategori, error)
	CreateCategori(input models.Kategoriinput) (models.Kategori, error)
	KelompokbykategoriID(ID string) ([]models.Kelompok, error)
	DeleteBarang(id string) error
	BulkDeleteBarang(ids []string) error
	UpdateBarang(input models.Barang) (models.Barang, error)
	// Findallhistory() ([]models.History, error)
	// Gethistorybymemberid(memberID string) ([]models.History, error)
	// Createhistory(input models.Historyinput) (models.History, error)
}

type barangService struct {
	repository repository.RepositoryBarang
}

func NewServiceBarang(repo repository.RepositoryBarang) Barangservice {
	return &barangService{repository: repo}
}

func (h *barangService) Getallbarang() ([]models.Barang, error) {
	return h.repository.Getallbarang()
}

func (h *barangService) Getallkategori() ([]models.Kategori, error) {
	return h.repository.Getallkategori()
}

func (h *barangService) Getbarangbyid(ID string) (models.Barang, error) {
	return h.repository.GetbarangByID(ID)
}
func (h *barangService) KelompokbykategoriID(ID string) ([]models.Kelompok, error) {
	return h.repository.KelompokbykategoriID(ID)
}

func (h *barangService) Createbarang(input models.Baranginput) (models.Barang, error) {

	uid := strings.ToUpper(uuid.New().String())
	newhistory := models.Barang{
		Harga:           input.Harga,
		Stok:            input.Stok,
		Id:              uid,
		Nama_barang:     input.Nama_barang,
		Kategori_id:     input.Kategori_id,
		Kelompok_barang: input.Kelompok,
	}
	createhistory, err := h.repository.Createbarang(newhistory)
	if err != nil {
		log.Printf("Ada kesalahan saat membuat barang: %v", err)
	}
	return createhistory, err
}

func (h *barangService) CreateCategori(input models.Kategoriinput) (models.Kategori, error) {

	uid := strings.ToUpper(uuid.New().String())
	newhistory := models.Kategori{
		ID:            uid,
		Nama_kategori: input.Namakategori,
	}
	createhistory, err := h.repository.Createcategpri(newhistory)
	if err != nil {
		log.Printf("Ada kesalahan saat membuat barang: %v", err)
	}
	return createhistory, err
}

func (h *barangService) DeleteBarang(id string) error {
	err := h.repository.DeletebarangbyID(id)
	if err != nil {
		log.Printf("Ada kesalahan saat menghapus barang: %v", err)
	}
	return err
}

func (s *barangService) BulkDeleteBarang(ids []string) error {
	return s.repository.BulkDeleteByIDs(ids)
}

func (s *barangService) UpdateBarang(input models.Barang) (models.Barang, error) {
	updateData := models.Barang{
		Id:              input.Id,
		Harga:           input.Harga,
		Stok:            input.Stok,
		Nama_barang:     input.Nama_barang,
		Kategori_id:     input.Kategori_id,
		Kelompok_barang: input.Kelompok_barang,
	}

	updatedBarang, err := s.repository.UpdatebarangID(updateData)
	if err != nil {
		return models.Barang{}, err
	}

	return updatedBarang, nil
}
