package handler

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Historyhandler struct {
	service service.Historyservice
}

func NewHistoryHistory(service service.Historyservice) *Historyhandler {
	return &Historyhandler{service: service}
}

func (h *Historyhandler) CreateHistory(c *gin.Context) {
	var input models.Historyinput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	history, err := h.service.Createhistory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Historydata": history,
	})
}

func (h *Historyhandler) GetHistoryByMemberID(c *gin.Context) {
	var input struct {
		MemberID string `json:"member_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	histories, err := h.service.Gethistorybymemberid(input.MemberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"histories": histories})
}
