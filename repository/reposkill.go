package repository

import (
	models "gin-gonic-gorm/model"
	"log"

	"gorm.io/gorm"
)

type Skillrepository interface {
	Createskill(skill models.Skill) (models.Skill, error)
	Findallskill() ([]models.Skill, error)
	Getskillbymemberid(ID string) (models.Skill, error)
	FindBySkillname(skillname string) (models.Skill, error)
}

type repositoryskil struct {
	db *gorm.DB
}

func NewRepositoryskill(db *gorm.DB) Skillrepository {
	return &repositoryskil{db: db}
}

func (r *repositoryskil) Createskill(skill models.Skill) (models.Skill, error) {
	err := r.db.Create(&skill).Error
	if err != nil {
		log.Printf("Gagal menyimpan skill: %v", err)
	}
	return skill, err
}

func (r *repositoryskil) Findallskill() ([]models.Skill, error) {
	var Skill []models.Skill
	err := r.db.Find(&Skill).Error
	return Skill, err
}

func (r *repositoryskil) Getskillbymemberid(ID string) (models.Skill, error) {
	var skill models.Skill
	err := r.db.Where("ID = ?", ID).Find(&skill).Error
	return skill, err
}
func (r *repositoryskil) FindBySkillname(skillname string) (models.Skill, error) {
	var skill models.Skill
	err := r.db.Where("skillname = ?", skillname).First(&skill).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return skill, err
	}
	return skill, nil
}
