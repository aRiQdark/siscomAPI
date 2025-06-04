package repository

import (
	models "gin-gonic-gorm/model"
	"log"

	"gorm.io/gorm"
)

type Historyrepository interface {
	Findallhistory() ([]models.History, error)
	Gethistorybymemberid(memberID string) ([]models.History, error)
	Createhistory(history models.History) (models.History, error)
}

type repositoryhistory struct {
	db *gorm.DB
}

func NewRepositoryhistory(db *gorm.DB) Historyrepository {
	return &repositoryhistory{db: db}
}

// Createhistory implements Historyrepository
func (r *repositoryhistory) Createhistory(history models.History) (models.History, error) {
	err := r.db.Create(&history).Error
	if err != nil {
		log.Printf("Error saving history: %v", err)
	}
	return history, err
}

// Findallhistory implements Historyrepository
func (r *repositoryhistory) Findallhistory() ([]models.History, error) {
	var history []models.History
	err := r.db.Find(&history).Error
	return history, err
}

// Gethistorybymemberid implements Historyrepository
func (r *repositoryhistory) Gethistorybymemberid(memberID string) ([]models.History, error) {
	var history []models.History
	err := r.db.Where("member_id = ?", memberID).Order("created_at desc").Find(&history).Error
	return history, err
}
