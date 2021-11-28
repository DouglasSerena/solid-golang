package repositories

import (
	"log"

	"github.com/DouglasSerena/solid-go/src/domain"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindBy(id string) (*domain.User, error)
	Update(user *domain.User, id string) (*domain.User, error)
	Delete(id string) error
}

type UserRepositoryDb struct {
	IUserRepository
	Db *gorm.DB
}

func (repo *UserRepositoryDb) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := repo.Db.Find(&users).Error

	if err != nil {
		log.Fatalf("Erro durante a criação do usuario: %v", err)
		return nil, err
	}

	return users, nil
}

func (repo *UserRepositoryDb) FindBy(id string) (*domain.User, error) {
	user := domain.User{}
	err := repo.Db.Find(&user, "id = ?", id).Error

	if err != nil {
		log.Fatalf("Erro durante a criação do usuario: %v", err)
		return &user, err
	}

	return &user, nil
}

func (repo *UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepere()

	if err != nil {
		log.Fatalf("Erro durante a validação do usuario: %v", err)
		return nil, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Erro durante a criação do usuario: %v", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepositoryDb) Update(user *domain.User, id string) (*domain.User, error) {
	err := repo.Db.Update(&user, "id = ?", id).Error

	if err != nil {
		log.Fatalf("Erro durante a atualização do usuario: %v", err)
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryDb) Delete(id string) error {
	err := repo.Db.Delete(&domain.User{}, "id = ?", id).Error

	if err != nil {
		log.Fatalf("Erro durante a criação do usuario: %v", err)
		return err
	}

	return nil
}
