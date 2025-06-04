package controller

import (
	// models "gin-gonic-gorm/model"
	"gin-gonic-gorm/service"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

type MemberController struct {
	service service.MemberService
}

func NewMemberController(service service.MemberService) *MemberController {
	return &MemberController{service: service}
}

// func (c *MemberController) Login(ctx *gin.Context) {
// 	var input models.LoginInput
// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token,, err := c.service.Login(input)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"token": token})
// }
