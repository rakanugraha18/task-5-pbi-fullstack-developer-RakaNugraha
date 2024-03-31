package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key;auto_increment; unique" json:"id"`
	Username  string    `gorm:"size:255;not null;" json:"username"`
	Email     string    `gorm:"size:255;not null; unique" json:"email"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"photos"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Photo struct {
	ID       int    `gorm:"primary_key;auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photo_url"`
	UserID   int    `gorm:"type:int" json:"user_id"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
