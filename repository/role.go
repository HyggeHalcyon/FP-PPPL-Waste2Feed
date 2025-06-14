package repository

import (
	"Waste2Feed/entity"

	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		GetbyId(string) (entity.Role, error)
		GetByName(string) (entity.Role, error)
	}

	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetbyId(roleId string) (entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("id = ?", roleId).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) GetByName(roleName string) (entity.Role, error) {
	var role entity.Role
	if err := r.db.Where("name = ?", roleName).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}
	return role, nil
}
