package service

import (
	"crypto/tls"
	"errors"
	"fmt"
	util "gin-gonic-gorm/Utils"
	"gin-gonic-gorm/entity"
	models "gin-gonic-gorm/model"
	"gin-gonic-gorm/repository"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	// "unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type MemberService interface {
	FindAllMembers() ([]entity.Membersdetail, error)
	FindByID(memberID string) (entity.Membersdetail, error)
	Create(input models.MemberInput) (entity.Membersdetail, error)
	Login(input models.LoginInput) (string, entity.Membersdetail, error)
	VerifyEmail(email string, kodeVerif string) (entity.Membersdetail, error)
	UpdatePassword(email, newPassword string) (entity.Membersdetail, error)
	FindAllMembersworker() ([]entity.Membersdetail, error)
	BulkVerifyEmail(email string) (entity.Membersdetail, error)
	OTPverifforgetpassword(email string, kodeVerif string) (entity.Membersdetail, error)
	UpdateImeimember(data entity.Membersdetail) (entity.Membersdetail, error)
	UpdateLongtitudelatitude(data entity.Membersdetail) (entity.Membersdetail, error)
	GetNearbyMembers(lat, lon, radius float64) ([]entity.Membersdetail, error)
}

type memberService struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) MemberService {
	return &memberService{repository: repo}
}

func (s *memberService) FindAllMembers() ([]entity.Membersdetail, error) {
	return s.repository.FindAllMembers()
}

func (s *memberService) FindAllMembersworker() ([]entity.Membersdetail, error) {
	return s.repository.FindAllMembersworker()
}
func (s *memberService) FindByID(memberID string) (entity.Membersdetail, error) {
	return s.repository.FindByID(memberID)
}

func (s *memberService) Create(input models.MemberInput) (entity.Membersdetail, error) {

	existingMember, err := s.repository.FindByUsername(input.Username)
	if err == nil && existingMember.Username != "" {
		return entity.Membersdetail{}, errors.New("username sudah terdaftar")
	}
	existingEmail, err := s.repository.FindByEmail(input.Email)
	if err == nil && existingEmail.Email != "" {
		return entity.Membersdetail{}, errors.New("email sudah terdaftar")
	}
	// if input.Kecamatan == "" {
	// 	return entity.Membersdetail{}, errors.New("Kecamatan harus diisi")
	// }
	existinghandphone, err := s.repository.FindByhandphone(input.Handphone)
	if err == nil && existinghandphone.Handphone != "" {
		return entity.Membersdetail{}, errors.New("Nomor handphone sudah terdaftar")
	}
	if input.Password == "" {
		return entity.Membersdetail{}, errors.New("Password harus diisi")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Membersdetail{}, errors.New("Gagal bycript password")
	}
	if input.Email == "" {
		return entity.Membersdetail{}, errors.New("Email harus diisi")
	}
	if !strings.Contains(input.Email, "@") {
		return entity.Membersdetail{}, errors.New("Email tidak sesuai format")
	}
	otp := generateOTP()
	uid := strings.ToUpper(uuid.New().String())
	newMember := entity.Membersdetail{
		MemberID:      uid,
		Fullname:      input.Fullname,
		Username:      input.Username,
		Email:         input.Email,
		KodeVerif:     otp,
		Password:      string(hashedPassword),
		Handphone:     input.Handphone,
		Umur:          input.Umur,
		Alamat:        input.Alamat,
		Kota:          input.Kota,
		Provinsi:      input.Provinsi,
		Kelurahan:     input.Kelurahan,
		Kodepos:       input.Kodepos,
		Lastjob:       input.Lastjob,
		IsLocked:      "0",
		Skill:         input.Skill,
		BirthDate:     input.Birthdate,
		Rating:        0,
		Emailverified: "0",
		Nationality:   "indonesia",
		Recordstatus:  "12",
		Limitbalance:  "0",
		Membertype:    "1",
		ProfileImage:  existingMember.ProfileImage,
		VAnumber:      "0",
		Longtitude:    0,
		Lattitude:     0,
		Imei:          input.Imei,
		Kecamatan:     input.Kecamatan,
	}

	err = sendEmail(input.Email, otp)
	if err != nil {
		log.Printf("Gagal mengirimkan kode otp: %v", err)
		return entity.Membersdetail{}, errors.New("Gagal mengirimkan kode otp: %v")
	}

	createdMember, err := s.repository.Create(newMember)
	if err != nil {
		log.Printf("Ada kesalahan saat daftar: %v", err)
	}

	return createdMember, err
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(999999-100000) + 100000
	return strconv.Itoa(otp)
}

var lastSent = make(map[string]time.Time)

const cooldownDuration = 30 * time.Second

func sendEmail(to string, otp string) error {

	if lastTime, found := lastSent[to]; found {
		if time.Since(lastTime) < cooldownDuration {
			return fmt.Errorf("Mohon tunggu beberapa saat untuk resend code")
		}
	}

	fromName := "Triniti App"
	fromEmail := "ariqmuhammad015@gmail.com"
	password := "waag gaze cjiz orkj "
	smtpServer := "smtp.gmail.com"
	port := 587

	htmlBody := fmt.Sprintf(`
        <html>
        <head>
            <style>
                body { font-family: Arial, sans-serif; }
                .container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; }
                .header { text-align: center; margin-bottom: 20px; }
                .footer { text-align: center; margin-top: 20px; font-size: 12px; color: #888; }
                .code { font-size: 24px; font-weight: bold; color: #333; }
            </style>
        </head>
        <body>
            <div class="container">
                <div class="header">
                    <h1>OTP Kode Anda</h1>
                </div>
                <p>Halo,</p>
                <p>Berikut adalah kode OTP Anda:</p>
                <p class="code">%s</p>
                <p>Silakan gunakan kode ini untuk melanjutkan proses.</p>
                <div class="footer">
                    <p>Terima kasih,</p>
                    <p>Powered by Triniti</p>
                </div>
            </div>
        </body>
        </html>
    `, otp)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, fromEmail))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Kode OTP Anda")
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtpServer, port, fromEmail, password)
	d.TLSConfig = &tls.Config{
		ServerName: smtpServer,
	}

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	lastSent[to] = time.Now()

	return nil
}

func (s *memberService) UpdateImeimember(data entity.Membersdetail) (entity.Membersdetail, error) {
	member, err := s.repository.UpdateImeibyemail(data)
	if err != nil {
		return entity.Membersdetail{}, err
	}
	return member, nil
}

func (s *memberService) UpdateLongtitudelatitude(data entity.Membersdetail) (entity.Membersdetail, error) {
	member, err := s.repository.UpdateLongtitudelatitude(data)
	if err != nil {
		return entity.Membersdetail{}, err
	}

	return member, nil
}

func (s *memberService) VerifyEmail(email string, kodeVerif string) (entity.Membersdetail, error) {
	member, err := s.repository.FindByEmail(email)
	if err != nil {
		return member, errors.New("email tidak ditemukan")
	}

	if member.KodeVerif != kodeVerif {
		return member, errors.New("kode verifikasi tidak sesuai")
	}

	member.Recordstatus = "2"
	member.Emailverified = "1"
	member, err = s.repository.Update(member)
	if err != nil {
		return member, errors.New("gagal memperbarui status member")
	}

	return member, nil
}

// func isStrongPassword(password string) bool {
// 	var (
// 		hasMinLen  = false
// 		hasUpper   = false
// 		hasLower   = false
// 		hasNumber  = false
// 		hasSpecial = false
// 	)

// 	if len(password) >= 8 {
// 		hasMinLen = true
// 	}

// 	for _, char := range password {
// 		switch {
// 		case unicode.IsUpper(char):
// 			hasUpper = true
// 		case unicode.IsLower(char):
// 			hasLower = true
// 		case unicode.IsDigit(char):
// 			hasNumber = true
// 		case unicode.IsPunct(char) || unicode.IsSymbol(char):
// 			hasSpecial = true
// 		}
// 	}

// 	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
// }

func isPasswordStartsWithUpperAndHasNumber(password string) bool {
	matched, _ := regexp.MatchString(`^[A-Z].*[0-9]`, password)
	return matched
}

func (s *memberService) UpdatePassword(email, newPassword string) (entity.Membersdetail, error) {
	if newPassword == "" {
		return entity.Membersdetail{}, errors.New("password tidak boleh kosong")
	}

	// if !isStrongPassword(newPassword) {
	// 	return entity.Membersdetail{}, errors.New("password harus mengandung alfabet dan karakter unik")
	// }

	if !isPasswordStartsWithUpperAndHasNumber(newPassword) {
		return entity.Membersdetail{}, errors.New("Password harus diawali dengan huruf besar dan mengandung angka")
	}

	existingMember, err := s.repository.FindByEmail(email)
	if err != nil {
		return entity.Membersdetail{}, errors.New("pengguna tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingMember.Password), []byte(newPassword))
	if err == nil {
		return entity.Membersdetail{}, errors.New("password tidak boleh sama dengan sebelumnya")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return entity.Membersdetail{}, errors.New("gagal mengenkripsi password")
	}

	existingMember.Password = string(hashedPassword)
	updatedMember, err := s.repository.Update(existingMember)
	if err != nil {
		return entity.Membersdetail{}, errors.New("gagal update password")
	}

	return updatedMember, nil
}

// bisa pakai resend kirim otp
func (s *memberService) BulkVerifyEmail(email string) (entity.Membersdetail, error) {
	member, err := s.repository.FindByEmail(email)
	if err != nil {
		return member, errors.New("email tidak ditemukan")
	}
	err = sendEmail(email, member.KodeVerif)
	if err != nil {
		return entity.Membersdetail{}, errors.New("Mohon tunggu beberapa saat untuk resend code")
	}
	return member, nil
}

func (s *memberService) OTPverifforgetpassword(email string, kodeVerif string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	member, err := s.repository.FindByEmail(email)
	if err != nil {
		return member, errors.New("email tidak ditemukan")
	}

	if member.KodeVerif != kodeVerif {
		return member, errors.New("kode verifikasi tidak sesuai")
	}

	return member, nil
}

// func (s *memberService) BulkUpdatePassword(email, newPassword string) (entity.Membersdetail, error) {
// 	if newPassword == "" {
// 		return entity.Membersdetail{}, errors.New("password tidak boleh kosong")
// 	}

// 	// if !isStrongPassword(newPassword) {
// 	// 	return entity.Membersdetail{}, errors.New("password harus mengandung alfabet dan karakter unik")
// 	// }

// 	if !isPasswordStartsWithUpperAndHasNumber(newPassword) {
// 		return entity.Membersdetail{}, errors.New("Password harus diawali dengan huruf besar dan mengandung angka")
// 	}

// 	existingMember, err := s.repository.FindByEmail(email)
// 	if err != nil {
// 		return entity.Membersdetail{}, errors.New("pengguna tidak ditemukan")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(existingMember.Password), []byte(newPassword))
// 	if err == nil {
// 		return entity.Membersdetail{}, errors.New("password tidak boleh sama dengan sebelumnya")
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		return entity.Membersdetail{}, errors.New("gagal mengenkripsi password")
// 	}

// 	existingMember.Password = string(hashedPassword)
// 	updatedMember, err := s.repository.Update(existingMember)
// 	if err != nil {
// 		return entity.Membersdetail{}, errors.New("gagal update password")
// 	}

// 	return updatedMember, nil
// }

func (s *memberService) Login(input models.LoginInput) (string, entity.Membersdetail, error) {
	var member entity.Membersdetail
	var err error

	if input.Email != "" {
		member, err = s.repository.FindByEmail(input.Email)
	} else if input.Handphone != "" {
		member, err = s.repository.FindByhandphone(input.Handphone)
	} else {
		return "", entity.Membersdetail{}, errors.New("Masukkan email atau Nomor handphone anda")
	}
	if err != nil {
		return "", entity.Membersdetail{}, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(input.Password))
	if err != nil {
		return "", entity.Membersdetail{}, errors.New("email atau password salah")
	}

	token, err := util.GenerateJWT(member.MemberID, member.Email)
	if err != nil {
		return "", entity.Membersdetail{}, errors.New("gagal mendapatkan token")
	}

	claims, err := util.ValidateJWT(token)
	if err != nil {
		return "", entity.Membersdetail{}, errors.New("gagal memvalidasi token")
	}

	decodedMemberID := member.MemberID
	decodedEmail := claims.Username

	fmt.Printf("Decoded MemberID: %s, Decoded Email: %s\n", decodedMemberID, decodedEmail)

	return token, member, nil
}

func (s *memberService) GetNearbyMembers(lat, lon, radius float64) ([]entity.Membersdetail, error) {
	return s.repository.FindNearbyMembers(lat, lon, radius)
}
