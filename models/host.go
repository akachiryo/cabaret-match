package models

import (
	"time"

	"gorm.io/gorm"
)

type GenderType int

const (
	Male GenderType = iota
	Female
	Other
)

var genderStrings = map[GenderType]string{
	Male:   "Male",
	Female: "Female",
	Other:  "Other",
}

func (g GenderType) String() string {
	return genderStrings[g]
}

type Host struct {
	ID                 uint           `gorm:"primaryKey"`
	Name               string         `gorm:"type:varchar(255)"`
	Gender             GenderType     `gorm:"type:int"`
	Place              string         `gorm:"type:varchar(255)"`
	Email              string         `gorm:"type:varchar(255);uniqueIndex"`
	Introduction       string         `gorm:"type:text"`
	Images             string         `gorm:"type:text"`
	Password           string         `gorm:"type:varchar(255)"`
	IsPresent          bool           `gorm:"default:false"`
	IsDiscard          bool           `gorm:"default:false"`
	ResetPasswordToken string         `gorm:"type:varchar(255);uniqueIndex"`
	ResetPasswordSentAt time.Time
	AllowPasswordChange bool          `gorm:"default:false"`
	ConfirmationToken   string        `gorm:"type:varchar(255);uniqueIndex"`
	ConfirmedAt         time.Time
	ConfirmationSentAt  time.Time
	UnconfirmedEmail    string        `gorm:"type:varchar(255)"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt 					gorm.DeletedAt `gorm:"index"`
}
