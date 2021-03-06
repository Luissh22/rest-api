package comment

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Slug      string         `json:"slug"`
	Body      string         `json:"body"`
	Author    string         `json:"author"`
}
