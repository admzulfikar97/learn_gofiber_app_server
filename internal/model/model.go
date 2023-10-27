package model

import (
	"github.com/google/uuid"
)

type Note struct {
	// gorm.Model
	// ID       int       `gorm:"int;not null" json:"id,omitempty"`
	ID        uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	UUID      uuid.UUID `gorm:"type:uuid"`
	Title     string    `gorm:"varchar(255);not null" json:"title"`
	SubTitle  string    `gorm:"varchar(255);not null" json:"sub_title"`
	Text      string    `gorm:"varchar(255);not null" json:"text"`
	UpdatedAt string    `json:"updated_at"`
	CreatedAt string    `json:"created_at"`
}
