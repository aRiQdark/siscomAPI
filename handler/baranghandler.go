package handler

import (
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarangHandler struct {
	service service.Barangservice
}

func NewBarangHandler(service service.Barangservice) *BarangHandler {
	return &BarangHandler{service: service}
}

func (h *BarangHandler) CreateBarang(c *gin.Context) {
	var input models.Baranginput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received input: %+v", input)

	member, err := h.service.Createbarang(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		if err.Error() == "Ada kesalahan saat input barang" {
			c.JSON(http.StatusConflict, gin.H{"error": "Ada kesalahan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	log.Printf("Created Barang: %+v", member)
	//inputrequestnya
	c.JSON(http.StatusOK, gin.H{
		"data": member,
	})

}

func (h *BarangHandler) FindBarangDetailById(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, _ := h.service.Getbarangbyid(input.ID)

	c.JSON(http.StatusOK, gin.H{"detailbarang": member})
}

func (h *BarangHandler) GetAllBarang(c *gin.Context) {
	members, err := h.service.Getallbarang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *BarangHandler) GetAllKategori(c *gin.Context) {
	members, err := h.service.Getallkategori()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *BarangHandler) CreateCategory(c *gin.Context) {
	var input models.Kategoriinput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received input: %+v", input)

	member, err := h.service.CreateCategori(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		if err.Error() == "Ada kesalahan saat input barang" {
			c.JSON(http.StatusConflict, gin.H{"error": "Ada kesalahan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	log.Printf("Created Barang: %+v", member)
	//inputrequestnya
	c.JSON(http.StatusOK, gin.H{
		"data": member,
	})

}

func (h *BarangHandler) FindKelompoklById(c *gin.Context) {
	var input struct {
		Idkategori string `json:"Idkategori" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, _ := h.service.KelompokbykategoriID(input.Idkategori)

	c.JSON(http.StatusOK, gin.H{"detailbarang": member})
}

func (h *BarangHandler) DeleteBarang(c *gin.Context) {
	id := c.Query("id") // Ambil dari query parameter, misalnya: /DeleteBarangbyid?id=123
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID barang wajib diisi"})
		return
	}

	err := h.service.DeleteBarang(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus barang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barang berhasil dihapus"})
}

func (h *BarangHandler) DeleteBarangBulk(c *gin.Context) {
	ids := c.QueryArray("ids") // ini akan menangkap multiple ids seperti ?ids=1&ids=2&ids=3

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	err := h.service.BulkDeleteBarang(ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus barang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}

func (h *BarangHandler) UpdateBarang(c *gin.Context) {
	var input models.Barang

	// Bind JSON dari request body ke struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
		return
	}

	// Pastikan ID barang tidak kosong
	if input.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID barang wajib diisi"})
		return
	}

	// Panggil service untuk update barang
	updatedBarang, err := h.service.UpdateBarang(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate barang: " + err.Error()})
		return
	}

	// Berhasil, kembalikan data barang yang sudah diupdate
	c.JSON(http.StatusOK, gin.H{"barang": updatedBarang})
}
