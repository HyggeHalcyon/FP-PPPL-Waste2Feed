package repository

import (
	"Waste2Feed/entity"

	"gorm.io/gorm"
)

type (
	FodderRepository interface {
		Create(entity.Fodder) (entity.Fodder, error)
		Update(entity.Fodder) (entity.Fodder, error)
		GetAllUnpaid(string) ([]entity.Fodder, error)
		GetByID(string) (entity.Fodder, error)
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

func (r *fodderRepository) Create(fodder entity.Fodder) (entity.Fodder, error) {
	if err := r.db.Create(&fodder).Error; err != nil {
		return entity.Fodder{}, err
	}

	return fodder, nil
}

func (r *fodderRepository) Update(fodder entity.Fodder) (entity.Fodder, error) {
	if err := r.db.Updates(&fodder).Error; err != nil {
		return entity.Fodder{}, err
	}

	return fodder, nil
}

func (r *fodderRepository) GetAllUnpaid(farmerID string) ([]entity.Fodder, error) {
	var fodders []entity.Fodder

	err := r.db.Where("farmer_id = ? AND paid = false", farmerID).Find(&fodders).Error
	if err != nil {
		return nil, err
	}

	return fodders, nil
}

func (r *fodderRepository) GetByID(id string) (entity.Fodder, error) {
	var fodder entity.Fodder

	err := r.db.Where("id = ?", id).First(&fodder).Error
	if err != nil {
		return entity.Fodder{}, err
	}

	return fodder, nil
}
