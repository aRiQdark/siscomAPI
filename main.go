package main

import (
	"fmt"
	"log"

	"gin-gonic-gorm/entity"
	"gin-gonic-gorm/handler"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"gin-gonic-gorm/route"
	"gin-gonic-gorm/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost port=5432 dbname=Koperasi user=postgres password=Ariq2001 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection DB")
	}
	Db.AutoMigrate(&entity.Membersdetail{})
	Db.AutoMigrate(&entity.Account{})
	Db.AutoMigrate(&models.History{})
	Db.AutoMigrate(&models.Skill{})
	Db.AutoMigrate(&models.Modelpesanan{})
	Db.AutoMigrate(&models.Transactionmodel{})
	Db.AutoMigrate(&models.Barang{})
	Db.AutoMigrate(&models.Kategori{})
	Db.AutoMigrate(&models.Kelompok{})

	AccountRepo := repository.NewRepositoryaccount(Db)
	fmt.Println(AccountRepo)
	Newrepo := repository.NewRepository(Db)
	Service := service.NewService(Newrepo)
	Historyrepo := repository.NewRepositoryhistory(Db)
	HistoryService := service.NewServiceHistory(Historyrepo)
	Historyhandler := handler.NewHistoryHistory(HistoryService)
	memberHandler := handler.NewMemberHandler(Service)
	Skillrepo := repository.NewRepositoryskill(Db)
	Skillservice := service.NewServiceSkill(Skillrepo)
	skillHandler := handler.NewSkill(Skillservice)
	Pesananrepo := repository.NewrepositoryPesanan(Db)
	Pesananservice := service.NewServicepesanan(Pesananrepo, Historyrepo, Newrepo)
	PesananHandler := handler.Newhandlerpesanan(Pesananservice)
	Transaksirepo := repository.NewRepository1(Db)
	TransaksiService := service.Newtransaction(Transaksirepo)
	Transaksihandler := handler.NewhandlerTransaksi(TransaksiService)
	BarangRepo := repository.NewRepositoryBarang(Db)
	BarangService := service.NewServiceBarang(BarangRepo)
	BarangHandler := handler.NewBarangHandler(BarangService)

	handlers := &handler.Handlers{
		MemberHandler:    memberHandler,
		HistoryHandler:   Historyhandler,
		Skillhandler:     skillHandler,
		Pesananhandler:   PesananHandler,
		Transaksihandler: Transaksihandler,
		BarangHandler:    BarangHandler,
	}

	r := gin.Default()
	route.Route(r, handlers)

	fmt.Println("DB connect sukses")

	if err := r.Run(); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
