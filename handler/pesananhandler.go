package handler

import (
	"encoding/base64"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"os"
	"unicode"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Pesananhandler struct {
	service service.PesananService
}

func Newhandlerpesanan(service service.PesananService) *Pesananhandler {
	return &Pesananhandler{
		service: service,
	}
}

func (h *Pesananhandler) Createpesanan(c *gin.Context) {
	var input models.Inputpesanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid: " + err.Error()})
		return
	}

	// Check if numeric fields are valid
	// if !isValidNumeric(input.Price) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Field Price tidak valid"})
	// 	return
	// }
	// if !isValidNumeric(input.Timestart) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Field Timestart tidak valid"})
	// 	return
	// }
	// if !isValidNumeric(input.Timeend) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Field Timeend tidak valid"})
	// 	return
	// }

	// Decode Base64 image
	imgData, err := base64.StdEncoding.DecodeString(input.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar tidak valid: " + err.Error()})
		return
	}

	// Create a temporary file to store the decoded image
	tmpFile, err := os.CreateTemp("", "upload-*.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat file sementara"})
		return
	}
	defer os.Remove(tmpFile.Name())

	// Write the decoded image data to the temporary file
	if _, err := tmpFile.Write(imgData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menulis file gambar"})
		return
	}
	tmpFile.Close()

	// Open the temporary file for reading
	imageFile, err := os.Open(tmpFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file gambar"})
		return
	}
	defer imageFile.Close()

	pesanan, err := h.service.Createpesanan(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membuat pesanan: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Pesanan anda berhasil",
		"createPesananData": pesanan,
	})
}

// Utility function to validate numeric fields
func isValidNumeric(value string) bool {
	for _, char := range value {
		if !unicode.IsDigit(char) && char != '.' {
			return false
		}
	}
	return true
}

func (h *Pesananhandler) GetpesananbymemberID(c *gin.Context) {
	var input models.Inputpesanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getpesananmemberID, err := h.service.Getpesananbymemberid(input.MemberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Pesanandata": getpesananmemberID,
	})
}
