package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	util "gin-gonic-gorm/Utils"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PesananService interface {
	Createpesanan(input models.Inputpesanan) (models.Modelpesanan, error)
	Getpesananbymemberid(memberID string) ([]models.Modelpesanan, error)
}

type pesananService struct {
	repository     repository.Pesananrepository
	historyService repository.Historyrepository
	Existingmember repository.Repository
	// contextTimeout time.Duration
}

// Constructor for PesananService
func NewServicepesanan(repo repository.Pesananrepository, historySvc repository.Historyrepository, member repository.Repository) PesananService {
	return &pesananService{
		Existingmember: member,
		// contextTimeout: time.Second * time.Duration(5000),
		repository:     repo,
		historyService: historySvc,
	}
}

func (h *pesananService) Createpesanan(input models.Inputpesanan) (models.Modelpesanan, error) {
	uid := strings.ToUpper(uuid.New().String())
	// ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	// defer cancel()
	// Validasi panjang MemberID
	if len(input.MemberID) < 34 {
		return models.Modelpesanan{}, errors.New("MemberID is too short")
	}

	trimmedUID := input.MemberID[34:]
	var tipe, nmpesanan, recordstatus string
	currentDatetime := time.Now().Format("20060102150405")

	// Set tipe dan nama berdasarkan input type
	switch input.Type {
	case "E0EDA025-0EB0-4981-A454-FE1401672BDA":
		tipe = "LTK"
		nmpesanan = "Listrik"
	case "063EBDC7-8241-4656-8944-C38A1967735E":
		tipe = "Air"
		nmpesanan = "Air"
	default:
		return models.Modelpesanan{}, errors.New("Tipe pesanan tidak dikenal")
	}
	recordstatus = "1"

	// Validasi input fields
	if input.Membername == "" {
		return models.Modelpesanan{}, errors.New("Nama pemesan tidak ditemukan")
	}
	if input.DestinationID == "" {
		return models.Modelpesanan{}, errors.New("Tujuan pesanan tidak diketahui")
	}
	// if input.Deskripsi == "" {
	// 	return models.Modelpesanan{}, errors.New("Deskripsi tidak boleh kosong")
	// }
	if input.Price == "" {
		return models.Modelpesanan{}, errors.New("Error price")
	}
	if input.Alamat == "" {
		return models.Modelpesanan{}, errors.New("Isi alamat tujuan")
	}
	if input.MemberID == "" {
		return models.Modelpesanan{}, errors.New("MemberID tidak ditemukan")
	}
	existingmember, _ := h.Existingmember.FindByID(input.MemberID)

	// Decode Base64 image
	imgData, err := base64.StdEncoding.DecodeString(input.Image)
	if err != nil {
		return models.Modelpesanan{}, errors.New("Gambar tidak valid: " + err.Error())
	}

	directory := "C:/foto kuli"
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return models.Modelpesanan{}, errors.New("Gagal membuat direktori untuk menyimpan gambar")
	}

	imagePath := fmt.Sprintf("%s/%s.png", directory, uid)

	err = os.WriteFile(imagePath, imgData, 0644)
	if err != nil {
		return models.Modelpesanan{}, errors.New("Gagal menyimpan file gambar")
	}
	log.Printf("Gambar disimpan di %s", imagePath)

	// Buat objek pesanan
	newPesanan := models.Modelpesanan{
		ID:              uid,
		MemberID:        input.MemberID,
		Type:            input.Type,
		Nopesanan:       tipe + trimmedUID + currentDatetime,
		Membername:      input.Membername,
		DestinationID:   input.DestinationID,
		DestinationName: input.DestinationName,
		SkillID:         input.SkillID,
		Skillname:       input.Skillname,
		Deskripsi:       input.Deskripsi,
		Timestart:       input.Timestart,
		Timeend:         input.Timeend,
		Price:           input.Price,
		Namapesanan:     nmpesanan,
		Recordstatus:    recordstatus,
		Kebutuhan:       " ",
		Sesipengerjaan:  input.Sesipengerjaan,
		Image:           imagePath,
		Alamat:          input.Alamat,
	}
	oneSignalMessageRaw := "Terima kasih telah mempercayakan NestFix, Pesanan atas nama " + existingmember.Username + " dengan kode pesanan " + newPesanan.Nopesanan + " akan di proses " + "."
	param := map[string]string{
		"id": "Kuli-LTK",
	}
	jsonString, _ := json.Marshal(param)
	if existingmember.Imei != "" {
		oneSignalMessage := util.RemoveHtmlTag(oneSignalMessageRaw)
		_, err := util.SendNotif(existingmember.Imei, "Nestfix "+" On progress", oneSignalMessage, string(jsonString))
		if err != nil {
			return models.Modelpesanan{}, err
		}
	}
	// Panggil repository untuk membuat pesanan

	createPesanan, err := h.repository.Createpesanan(newPesanan)
	if err != nil {
		return models.Modelpesanan{}, errors.New("Ada kesalahan saat membuat pesanan")
	}

	// Buat catatan riwayat
	historyInput := models.History{
		ID:           uuid.New().String(),
		MemberID:     input.MemberID,
		Price:        input.Price,
		Title:        input.Deskripsi,
		Historyname:  nmpesanan,
		Recordstatus: recordstatus,
	}
	if recordstatus == "1" {
		if _, err = h.historyService.Createhistory(historyInput); err != nil {
			log.Printf("Gagal menyimpan riwayat pesanan: %v", err)
			return models.Modelpesanan{}, errors.New("Gagal menyimpan riwayat pesanan")
		}
		log.Println("Riwayat pesanan disimpan dengan sukses.")
	}

	return createPesanan, nil
}

func (h *pesananService) Getpesananbymemberid(memberID string) ([]models.Modelpesanan, error) {
	return h.repository.FindpesananBymemberid(memberID)
}
