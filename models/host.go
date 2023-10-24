package models

import (
	"api/database"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	Password           string         `gorm:"size:255;not null;" json:"password"`
	IsPresent          bool           `gorm:"default:false"`
	IsDiscard          bool           `gorm:"default:false"`
	gorm.Model
	ResetPasswordToken string         `gorm:"type:varchar(255);uniqueIndex"`
	ResetPasswordSentAt *time.Time
	AllowPasswordChange bool          `gorm:"default:false"`
	ConfirmationToken   string        `gorm:"type:varchar(255);uniqueIndex"`
	ConfirmedAt         *time.Time
	ConfirmationSentAt  *time.Time
	UnconfirmedEmail    string        `gorm:"type:varchar(255)"`
}

type hostResponse struct {
	ID    uint
	Email string
}


func (h *Host) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(h.Password), bcrypt.DefaultCost)

	if err != nil {
			return err
	}

	h.Password = string(hashedPassword)

	h.Email = strings.ToLower(h.Email)

	return nil
}

func (h Host) Save() (hostResponse, error) {
	err := database.DB.Create(&h).Error

	if err != nil {
		return hostResponse{}, err
	}

	data := hostResponse{
		ID:    h.ID,
		Email: h.Email,
	}
	return data, nil
}


func HostAuthenticated(email string, password string) (hostResponse, error) {
	var host Host

	err := database.DB.Where("email = ?", email).First(&host).Error
	if err != nil {
		return hostResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(host.Password), []byte(password))
	if err != nil {
		return hostResponse{}, err
	}

	data := hostResponse{
		ID:    host.ID,
		Email: host.Email,
	}

	return data, nil
}
