package seeder

import (
	"Waste2Feed/migrations/seeder/seeders"

	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) error {
	if err := seeders.RoleSeeder(db); err != nil {
		return err
	}

	if err := seeders.UserSeeder(db); err != nil {
		return err
	}
	return nil
}
