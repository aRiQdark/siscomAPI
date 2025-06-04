package entity

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Membersdetail struct {
	MemberID      string
	Fullname      string
	Username      string
	Email         string
	Handphone     string
	Password      string
	Umur          string
	Alamat        string
	Kota          string
	Provinsi      string
	Kelurahan     string
	Kodepos       string
	Kecamatan     string
	Lastjob       pq.StringArray `gorm:"type:text[]"`
	Skill         pq.StringArray `gorm:"type:text[]"`
	Recordstatus  string
	KodeVerif     string
	Membertype    string
	IsLocked      string
	ProfileImage  sql.NullString
	BirthDate     string
	Nationality   string
	Rating        int32
	Emailverified string
	Limitbalance  string
	VAnumber      string
	Lattitude     float64 `json:"lattitude"`
	Longtitude    float64 `json:"longtitude"`
	Imei          string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
