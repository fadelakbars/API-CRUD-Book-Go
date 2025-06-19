package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uuid.UUID 	`gorm:"primaryKey;type:char(36)" json:"id"`
	Title  string    	`gorm:"column:title;type:varchar(255);not null" json:"title" validate:"required"`
	Email  string    	`gorm:"column:email;type:varchar(255);not null" json:"email" validate:"required"`
	Author string    	`gorm:"column:author;type:varchar(255);not null" json:"author" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	b.CreatedAt = time.Now()
	return
}

// test push satu file