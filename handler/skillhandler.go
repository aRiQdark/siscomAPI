package handler

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Skillhandler struct {
	service service.Skillservice
}

func NewSkill(service service.Skillservice) *Skillhandler {
	return &Skillhandler{service: service}
}

func (h *Skillhandler) Createskill(c *gin.Context) {
	var input models.Skillinput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.service.Createskill(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Skilldata": skill,
	})
}

func (h *Skillhandler) GetSkillByID(c *gin.Context) {
	var input struct {
		ID string `json:"ID" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.service.FindskillByID(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"skill": skill})
}

func (h *Skillhandler) GetAllSkill(c *gin.Context) {
	skill, err := h.service.Findallskill()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}
