package seeders

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"Waste2Feed/entity"

	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	jsonFile, err := os.Open("./migrations/seeder/json/user.json")
	if err != nil {
		return err
	}
	jsonData, _ := io.ReadAll(jsonFile)

	var listUser []entity.User
	json.Unmarshal(jsonData, &listUser)

	// only create if it does not exist
	for _, data := range listUser {
		var user entity.User
		if data.Email != "" {
			err := db.Where(&entity.User{Email: data.Email}).First(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			exist := db.Find(&user, "email = ?", data.Email).RowsAffected
			if exist == 0 {
				if err := db.Create(&data).Error; err != nil {
					return err
				}
			}
		} else if data.PhoneNumber != "" {
			err := db.Where(&entity.User{PhoneNumber: data.PhoneNumber}).First(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			exist := db.Find(&user, "phone_number = ?", data.PhoneNumber).RowsAffected
			if exist == 0 {
				if err := db.Create(&data).Error; err != nil {
					return err
				}
			}
		} else {
			return errors.New("user must have either email or phone number")
		}
	}

	return nil
}
