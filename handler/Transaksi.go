package handler

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaksihandel struct {
	service service.Transaksiservice
}

func NewhandlerTransaksi(service service.Transaksiservice) *Transaksihandel {
	return &Transaksihandel{service: service}
}

func (h *Transaksihandel) Createtransaksi(c *gin.Context) {
	var input models.Transaksiinput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaksi, err := h.service.Storetransaction(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message":           "Transaksi anda berhasil",
		"Createpesanandata": transaksi,
	})
}

func (h *Transaksihandel) Findalltransaksi(c *gin.Context) {
	transaksi, err := h.service.Readdatatransaksi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaksi)
}
