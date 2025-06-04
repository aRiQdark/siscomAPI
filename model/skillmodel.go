package models

import "time"

type Skill struct {
	ID             string
	Skillname      string
	DailyRate      string
	AfternoonHours string
	EveningHours   string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
