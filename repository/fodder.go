package repository

import (
	"gorm.io/gorm"
)

type (
	FodderRepository interface {
		// Create()
	}

	fodderRepository struct {
		db *gorm.DB
	}
)

func NewFodderRepository(db *gorm.DB) FodderRepository {
	return &fodderRepository{
		db: db,
	}
}
