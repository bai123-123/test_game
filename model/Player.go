package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Player struct {
	gorm.Model
	Username string  `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string  `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Email    string  `gorm:"type:varchar(500);not null" json:"email"`
	Score    float64 `gorm:"default 0" json:"score"`
}

func CreatePlayer(data *Player) error {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckExist(email string) bool {
	var player Player
	db.Where("email = ?", email).First(&player)
	if player.ID > 0 {
		return true
	}
	return false
}

func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

func CheckLogin(email string, password string) (Player, error) {
	var player Player
	var PasswordErr error

	db.Where("email = ?", email).First(&player)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(password))

	if player.ID == 0 {
		return player, PasswordErr
	}
	if PasswordErr != nil {
		return player, PasswordErr
	}
	return player, nil
}
