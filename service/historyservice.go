package service

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"log"
	"strings"

	"github.com/google/uuid"
)

type Historyservice interface {
	Findallhistory() ([]models.History, error)
	Gethistorybymemberid(memberID string) ([]models.History, error)
	Createhistory(input models.Historyinput) (models.History, error)
}

type historyService struct {
	repository repository.Historyrepository
}

func NewServiceHistory(repo repository.Historyrepository) Historyservice {
	return &historyService{repository: repo}
}

// Createhistory implements Historyservice
func (h *historyService) Createhistory(input models.Historyinput) (models.History, error) {
	uid := strings.ToUpper(uuid.New().String())
	newhistory := models.History{
		ID:           uid,
		MemberID:     input.MemberID,
		Price:        input.Price,
		Recordstatus: "1",
		Historyname:  input.Deskripsi,
		Title:        input.Title,
	}
	createhistory, err := h.repository.Createhistory(newhistory)
	if err != nil {
		log.Printf("Ada kesalahan saat daftar: %v", err)
	}
	return createhistory, err
}

func (h *historyService) Findallhistory() ([]models.History, error) {
	return h.repository.Findallhistory()
}

func (h *historyService) Gethistorybymemberid(memberID string) ([]models.History, error) {
	return h.repository.Gethistorybymemberid(memberID)
}

// func (h *memberService) FindByID(memberID string) (entity.Membersdetail, error) {
// 	return s.repository.FindByID(memberID)
// }
