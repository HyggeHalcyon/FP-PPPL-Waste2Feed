package repository

import (
	"Waste2Feed/entity"
	"math"

	"gorm.io/gorm"
)

type (
	PickupRepository interface {
		Create(entity.Pickup) (entity.Pickup, error)
		GetPaginationByUser(string, int, int) ([]entity.Pickup, int64, int64, error)
		GetPaginationByCourier(string, int, int) ([]entity.Pickup, int64, int64, error)
	}

	pickupRepository struct {
		db *gorm.DB
	}
)

func NewPickupRepository(db *gorm.DB) PickupRepository {
	return &pickupRepository{
		db: db,
	}
}

func (r *pickupRepository) Create(pickup entity.Pickup) (entity.Pickup, error) {
	if err := r.db.Create(&pickup).Error; err != nil {
		return entity.Pickup{}, err
	}

	return pickup, nil
}

func (r *pickupRepository) GetPaginationByUser(userID string, limit, page int) ([]entity.Pickup, int64, int64, error) {
	var files []entity.Pickup
	var count int64

	err := r.db.Where("user_id = ?", userID).Model(&entity.Pickup{}).Count(&count).Error
	if err != nil {
		return nil, 0, 0, err
	}

	maxPage := int64(math.Ceil(float64(count) / float64(limit)))
	offset := (page - 1) * limit

	err = r.db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&files).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return files, maxPage, count, nil
}

func (r *pickupRepository) GetPaginationByCourier(courierID string, limit, page int) ([]entity.Pickup, int64, int64, error) {
	var files []entity.Pickup
	var count int64

	err := r.db.Where("courier_id = ?", courierID).Model(&entity.Pickup{}).Count(&count).Error
	if err != nil {
		return nil, 0, 0, err
	}

	maxPage := int64(math.Ceil(float64(count) / float64(limit)))
	offset := (page - 1) * limit

	err = r.db.Where("courier_id = ?", courierID).Offset(offset).Limit(limit).Find(&files).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return files, maxPage, count, nil
}
