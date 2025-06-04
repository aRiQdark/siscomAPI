package models

import "time"

type History struct {
	MemberID     string
	ID           string
	Title        string
	Price        string
	Recordstatus string
	Historyname  string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
