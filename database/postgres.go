package database

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/Devansh3712/go-banking-api/config"
	"github.com/Devansh3712/go-banking-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PostgreSQL URI of the database used by the API.
var databaseURI string = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s",
	config.GetEnv("POSTGRES_HOSTNAME"),
	config.GetEnv("POSTGRES_USERNAME"),
	config.GetEnv("POSTGRES_PASSWORD"),
	config.GetEnv("POSTGRES_DATABASE"),
)

// Hash passwords using SHA256.
func hash(password string) (*string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return &result, nil
}

// Create tables using structs.
func MigratePostgres() error {
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
	passwordHash, err := hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = *passwordHash
	accNumber, _ := rand.Prime(rand.Reader, 32)
	account := &models.Account{
		Email:   user.Email,
		User:    user,
		Balance: 0,
		Number:  accNumber.String(),
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
	data, err := hash(user.Password)
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
func GetAccountData(email string) (*models.Account, error) {
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

// Update the balance of a user's account.
func UpdateAccountBalance(email string, balance uint64) error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}
	if query := db.Model(&models.Account{}).Where("email = ?", email).Update("balance", balance); query.Error != nil {
		return query.Error
	}
	return nil
}
