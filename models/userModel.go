package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	First_Name  string    `gorm:"not null;type:varchar(50);default:null"`
	Middle_Name string    `gorm:"not null;type:varchar(50);default:null"`
	Last_Name   string    `gorm:"not null;type:varchar(50);default:null"`
	Email       string    `gorm:"unique;not null;type:varchar(100);default:null"`
	Phonenumber string    `gorm:"unique;not null;type:varchar(50);default:null"`
	Password    string    `gorm:"not null;type:varchar(100);default:null"`
	Avatar      string    `gorm:"default:''"`
	IsAdmin     bool      `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func (u *User) BeforeSave(tx *gorm.DB) error {
	err := u.HashPassword()
	if err != nil {
		return err
	}
	return nil
}
