package route

import (
	// "gin-gonic-gorm/controller"
	"gin-gonic-gorm/handler"
	"gin-gonic-gorm/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	MemberHandler  *handler.MemberHandler
	HistoryHandler *handler.Historyhandler
	Skillhandler   *handler.Skillhandler
	Pesananhandler *handler.Pesananhandler
	// Tambahkan handler lain di sini
}

func Route(app *gin.Engine, handlers *handler.Handlers) {
	route := app
	auth := route.Group("/")
	route.GET("/")
	route.POST("/users", handlers.MemberHandler.CreateMember)
	// route.PUT("/v1", controller.Members)
	route.POST("/login", handlers.MemberHandler.Login)
	route.POST("/createhistory", handlers.HistoryHandler.CreateHistory)
	route.GET("/allmembersworker", handlers.MemberHandler.Getallmemberworker)
	route.POST("/verifyemail", handlers.MemberHandler.VerifyEmailByOTP)
	route.POST("/BulkVerifyEmail", handlers.MemberHandler.BulkVerifyEmail)
	route.POST("/OTPverifforgetpasssword", handlers.MemberHandler.OTPverifforgetpasssword)
	route.GET("/allmember", handlers.MemberHandler.GetAllMembers)
	route.GET("/skill", handlers.Skillhandler.GetAllSkill)
	route.POST("/createskill", handlers.Skillhandler.Createskill)
	route.POST("/Getskillbyid", handlers.Skillhandler.GetSkillByID)
	route.POST("/Createtransaksi", handlers.Transaksihandler.Createtransaksi)
	route.POST("/updateimei", handlers.MemberHandler.Updateimeibyemail)
	route.GET("/Finddata", handlers.Transaksihandler.Findalltransaksi)
	route.POST("/updategeolocation", handlers.MemberHandler.UpdateLongtitudelatitude)
	route.GET("/GetNearbyMembers", handlers.MemberHandler.GetNearbyMembers)
	route.GET("/Getallbarang", handlers.BarangHandler.GetAllBarang)
	route.POST("/Createbarang", handlers.BarangHandler.CreateBarang)
	route.POST("/GetbarangbyID", handlers.BarangHandler.FindBarangDetailById)
	route.GET("/Getallkategori", handlers.BarangHandler.GetAllKategori)
	route.POST("/CreateKategori", handlers.BarangHandler.CreateCategory)
	route.POST("/Findkategoribyid", handlers.BarangHandler.FindKelompoklById)
	route.POST("/Findkelompokbyid", handlers.BarangHandler.FindKelompoklById)
	route.DELETE("/DeleteBarangbyid", handlers.BarangHandler.DeleteBarang)
	route.DELETE("/DeleteBarangbulking", handlers.BarangHandler.DeleteBarangBulk)
	route.PUT("/UpdatebarangbyID", handlers.BarangHandler.UpdateBarang)

	//JWT
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/members/details", handlers.MemberHandler.FindMemberDetailById)
	auth.POST("/history", handlers.HistoryHandler.GetHistoryByMemberID)
	auth.POST("/updatepassword", handlers.MemberHandler.UpdatePassword)
	auth.POST("/Createpesanan", handlers.Pesananhandler.Createpesanan)
	auth.POST("/GetpesananbymemberID", handlers.Pesananhandler.GetpesananbymemberID)
}
