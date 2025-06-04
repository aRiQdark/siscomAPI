package controller

import (
	"gin-gonic-gorm/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Customer struct {
	Nama      string
	Handphone int
	isprimary string
}

type memberinput struct {
	membername string
	password   int
}

func Members(ctx *gin.Context) {
	members := entity.Membersdetail{}
	var membersinput memberinput
	uid := uuid.New()

	// if membersinput.membername == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"errors": "Nama member tidak boleh kosong",
	// 	})
	// 	return
	// }

	ctx.JSON(200, gin.H{
		members.Username: membersinput.membername,
		members.Password: membersinput.password,
		members.MemberID: uid,
	})

}

func Getalluser(ctx *gin.Context) {
	members := []string{"Ariq", "Budi", "Charlie"}
	member := "Ariq"
	accounts := map[string]int{
		"simpanan deposito":  200000,
		"simpanan berjangka": 300000,
		"simpanan pokok":     100000,
	}
	isValidate := true

	isMember := false
	for _, m := range members {
		if m == member {
			isMember = true
			break
		}
	}
	if !isMember {
		ctx.AbortWithStatusJSON(400, gin.H{
			"Bukan": "bukan member",
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"isTrue": member,
		})
	}
	if !isValidate {
		ctx.AbortWithStatusJSON(200, gin.H{
			"Bad": "Request",
		})
		return
	}

	primaryAccount := ""
	for acc := range accounts {
		if acc == "simpanan berjangka" {
			primaryAccount = acc
			break
		}
	}

	balance := accounts[primaryAccount]

	ctx.JSON(200, gin.H{
		"Username":        member,
		"accounts":        accounts,
		"primary_account": primaryAccount,
		"Saldo Balance":   balance,
	})
}

func Getusers(ctx *gin.Context) {
	var customerinput Customer
	err := ctx.ShouldBindJSON(&customerinput)
	if err != nil {
		log.Fatal(err)
	}
	accounts := map[string]int{
		"simpanan deposito":  200000,
		"simpanan berjangka": 300000,
		"simpanan pokok":     100000,
	}
	// primaryAccount := ""
	// for acc := range accounts {
	// 	if acc == "simpanan berjangka" {
	// 		primaryAccount = acc
	// 		break
	// 	}
	// }
	accountvalue, exist := accounts[customerinput.isprimary]
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "simpanan tidak ditemukan",
		})
	}
	if customerinput.Nama != "ariq" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bukan ariq",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"nama":      customerinput.Nama,
		"handphone": customerinput.Handphone,
		"isprimary": accountvalue,
	})
}
