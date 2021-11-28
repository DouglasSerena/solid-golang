package domain

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	Id       string    `json:"id" gorm:"type:uuid;primary_key"`
	CreateAt time.Time `json:"createAt" gorm:"type:datetime"`
	UpdateAt time.Time `json:"updateAt" gorm:"type:datetime"`
	DeleteAt time.Time `json:"deleteAt" gorm:"type:datetime" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreateAt", time.Now())

	if err != nil {
		log.Fatalf("Erro durante a geração do createAt: %v", err)
	}

	err = scope.SetColumn("Id", uuid.New().String())

	if err != nil {
		log.Fatalf("Erro durante a geração do id: %v", err)
	}

	return nil
}
