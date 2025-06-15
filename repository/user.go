package repository

import (
	"Waste2Feed/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(entity.User) (entity.User, error)
		Update(entity.User) (entity.User, error)
		Delete(entity.User) error
		GetAllByRoleID(string) ([]entity.User, error)
		GetById(string) (entity.User, error)
		GetByEmail(string) (entity.User, error)
		GetByPhoneNumber(string) (entity.User, error)
		GetByUsername(string) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllByRoleID(roleID string) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Where("role_id = ?", roleID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
	if err := r.db.Updates(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetById(userID string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByPhoneNumber(phoneNumber string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("phone_number = ?", phoneNumber).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("user_name = ?", username).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) Delete(user entity.User) error {
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
