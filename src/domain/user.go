package domain

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

func (user *User) NewUser() *User {
	return &User{}
}

func (user *User) Prepere() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Erro durante a geração do password: %v", err)
		return err
	}

	user.Password = string(password)
	user.Token = uuid.New().String()

	err = user.validate()

	if err != nil {
		log.Fatalf("Erro na validação dos dados: %v", err)
		return err
	}

	return nil
}

func (user *User) validate() error {
	return nil
}
