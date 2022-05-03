package models

type Account struct {
	Email  string  `json:"email" gorm:"primaryKey"`
	User   *User   `json:"user" gorm:"foreignKey:Email"`
	Amount float32 `json:"amount" gorm:"default:0"`
}
