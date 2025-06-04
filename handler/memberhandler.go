package handler

import (
	"gin-gonic-gorm/entity"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	service service.MemberService
}

func NewMemberHandler(service service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) CreateMember(c *gin.Context) {
	var input models.MemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received input: %+v", input)

	member, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	log.Printf("Created member: %+v", member)
	//inputrequestnya
	c.JSON(http.StatusOK, gin.H{
		"data": member,
	})

}

func (h *MemberHandler) FindMemberDetailById(c *gin.Context) {
	var input struct {
		MemberID string `json:"member_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.service.FindByID(input.MemberID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"detailmember": member})
}

func (h *MemberHandler) VerifyEmailByOTP(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required,email"`
		KodeVerif string `json:"kode_verif" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.service.VerifyEmail(input.Email, input.KodeVerif)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email berhasil diverifikasi",
		"member":  member,
	})
}

func (h *MemberHandler) UpdatePassword(c *gin.Context) {
	var input models.MemberInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.UpdatePassword(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil di ubah",
	})
}

type LoginResponse struct {
	StatusCode int                    `json:"statuscode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

func (h *MemberHandler) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received login input: %+v", input)

	token, member, err := h.service.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Generated token: %s", token)

	response := LoginResponse{
		StatusCode: http.StatusOK,
		Message:    "Login berhasil",
		Data: map[string]interface{}{
			"id":           member.MemberID,
			"username":     member.Username,
			"handphone":    member.Handphone,
			"token":        token,
			"recordstatus": member.Recordstatus,
			"membertype":   member.Membertype,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *MemberHandler) GetAllMembers(c *gin.Context) {
	members, err := h.service.FindAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *MemberHandler) Getallmemberworker(c *gin.Context) {
	member, err := h.service.FindAllMembersworker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, member)
}
func (h *MemberHandler) BulkVerifyEmail(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.service.BulkVerifyEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kode berhasil di kirim",
		"Kode":    member.KodeVerif,
	})
}

func (h *MemberHandler) Updateimeibyemail(c *gin.Context) {
	var input entity.Membersdetail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	updatedMember, err := h.service.UpdateImeimember(input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update IMEI: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "IMEI updated successfully",
		"member":  updatedMember,
	})
}

func (h *MemberHandler) UpdateLongtitudelatitude(c *gin.Context) {
	var input entity.Membersdetail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	updatedMember, err := h.service.UpdateLongtitudelatitude(input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update Location: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Your location has been update",
		"member":  updatedMember,
	})
}

func (h *MemberHandler) OTPverifforgetpasssword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.OTPverifforgetpassword(input.Email, input.OTP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil ",
	})
}

func (h *MemberHandler) GetNearbyMembers(c *gin.Context) {
	latParam := c.Query("lat")
	lonParam := c.Query("lon")

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}

	lon, err := strconv.ParseFloat(lonParam, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	const radius = 0.2

	members, err := h.service.GetNearbyMembers(lat, lon, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nearby members"})
		return
	}

	// Hitung jumlah members, lalu kurangi 1 (pastikan gak jadi negatif)
	count := len(members)

	c.JSON(http.StatusOK, gin.H{
		"members": members,
		"count":   count,
	})
}

