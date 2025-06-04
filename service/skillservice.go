package service

import (
	"errors"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"strings"

	"github.com/google/uuid"
)

type Skillservice interface {
	Createskill(input models.Skillinput) (models.Skill, error)
	Findallskill() ([]models.Skill, error)
	FindskillByID(skillID string) (models.Skill, error)
	// FindskillByname(skillname string) ([]models.Skill, error)
}

type skillService struct {
	repository repository.Skillrepository
}

func NewServiceSkill(repo repository.Skillrepository) Skillservice {
	return &skillService{repository: repo}
}

func (h *skillService) Createskill(input models.Skillinput) (models.Skill, error) {
	existingSkill, err := h.repository.FindBySkillname(input.Skillname)
	if err != nil {
		return models.Skill{}, err
	}
	if existingSkill.ID != "" {
		return models.Skill{}, errors.New("Skill sudah ada sebelumnya")
	}

	uid := strings.ToUpper(uuid.New().String())
	newskill := models.Skill{
		ID:             uid,
		DailyRate:      input.DailyRate,
		AfternoonHours: input.AfternoonHours,
		EveningHours:   input.EveningHours,
		Skillname:      input.Skillname,
	}

	createskill, err := h.repository.Createskill(newskill)
	if err != nil {
		return models.Skill{}, errors.New("Ada kesalahan saat menghubungkan")
	}
	return createskill, err
}

func (h *skillService) Findallskill() ([]models.Skill, error) {
	return h.repository.Findallskill()
}

func (h *skillService) FindskillByID(skillID string) (models.Skill, error) {
	return h.repository.Getskillbymemberid(skillID)
}

// func (h *skillService) FindskillByname(skillname string) ([]models.Skill, error) {
// 	return h.repository.FindBySkillname(skillname)
// }
