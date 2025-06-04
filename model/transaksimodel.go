package models

import "time"

type Transactionmodel struct {
	TransactionID string
	Userid        string
	Contractid    string
	Nama          string
	Nominal       string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
