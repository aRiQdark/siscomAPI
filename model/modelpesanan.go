package models

import (
	"time"
)

type Modelpesanan struct {
	ID              string
	Namapesanan     string
	Type            string
	MemberID        string
	Nopesanan       string
	Membername      string
	DestinationID   string
	DestinationName string
	SkillID         string
	Skillname       string
	Deskripsi       string
	Timestart       string
	Timeend         string
	Price           string
	Kebutuhan       string
	Sesipengerjaan  string
	Image           string
	Recordstatus    string
	Alamat          string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
