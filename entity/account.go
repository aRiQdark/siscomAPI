package entity

import "time"

type Account struct {
	accoung_id     string
	memberid       string
	balance        string
	closingbalance string
	accounttype    string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
