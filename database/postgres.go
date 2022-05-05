package database

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/Devansh3712/go-banking-api/models"
	"github.com/Devansh3712/go-banking-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PostgreSQL URI of the database used by the API.
var databaseURI string = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s",
	config.EnvValue("HOST"), config.EnvValue("USER"),
	config.EnvValue("PASSWORD"), config.EnvValue("DBNAME"),
)

// Create tables using structs.
func Migrate() error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Account{}, &models.User{})
	return nil
}

// Add a new user to the database. The user password is hashed
// using SHA256.
func CreateUser(user *models.User) (*string, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	passwordHash, err := utils.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = *passwordHash
	accNumber, _ := rand.Prime(rand.Reader, 32)
	account := &models.Account{
		Email:  user.Email,
		User:   user,
		Amount: 0,
		Number: accNumber.String(),
	}
	if query := db.Create(account); query.Error != nil {
		return nil, query.Error
	}
	return &account.Number, nil
}

// Validate user password. The hash of password is compared
// in order to validate the password.
func AuthUser(user *models.UserAuth) error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}
	result := models.User{}
	if query := db.Where("email = ?", user.Email).First(&result); query.Error != nil {
		return query.Error
	}
	data, err := utils.Hash(user.Password)
	if err != nil {
		return err
	}
	if result.Password != *data {
		return errors.New("incorrect password")
	}
	return nil
}

// Return user data.
func GetUserData(email string) (*models.User, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := models.User{}
	if query := db.Where("email = ?", email).First(&result); query.Error != nil {
		return nil, query.Error
	}
	return &result, nil
}

// Return user's account data.
func GetUserAccountData(email string) (*models.Account, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := models.Account{}
	if query := db.Preload(clause.Associations).Where("email = ?", email).First(&result); query.Error != nil {
		return nil, query.Error
	}
	return &result, nil
}
