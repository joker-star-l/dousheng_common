package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        int64          `json:"id" gorm:"column:id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}
