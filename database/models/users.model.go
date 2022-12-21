package models

import "time"

type Users struct {
	ID        uint      `json:"id" gorm:"autoIncrement" gorm:"primaryKey: type:uuid"`
	Login     string    `json:"login" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
