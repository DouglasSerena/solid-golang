package usecase

import (
	"log"

	"github.com/DouglasSerena/solid-go/src/app/repositories"
	"github.com/DouglasSerena/solid-go/src/domain"
)

type IUserUsecase interface {
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User, id string) (*domain.User, error)
	FindBy(id string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Delete(id string) error
}

type UserUsecase struct {
	IUserUsecase
	UserRepository repositories.IUserRepository
}

func (userUsecase *UserUsecase) FindAll() ([]domain.User, error) {
	users, err := userUsecase.UserRepository.FindAll()

	if err != nil {
		log.Fatalf("Erro ao buscar totos os usuarios: %v", err)
		return nil, err
	}

	return users, nil
}

func (userUsecase *UserUsecase) FindBy(id string) (*domain.User, error) {
	user, err := userUsecase.UserRepository.FindBy(id)

	if err != nil {
		log.Fatalf("Erro ao buscar o usuario: %v", err)
		return nil, err
	}

	return user, nil
}

func (userUsecase *UserUsecase) Create(user *domain.User) (*domain.User, error) {
	user, err := userUsecase.UserRepository.Insert(user)

	if err != nil {
		log.Fatalf("Erro ao criar o usuario: %v", err)
		return user, err
	}

	return user, nil
}

func (userUsecase *UserUsecase) Update(user *domain.User, id string) (*domain.User, error) {
	user, err := userUsecase.UserRepository.Update(user, id)

	if err != nil {
		log.Fatalf("Erro ao criar o usuario: %v", err)
		return user, err
	}

	return user, nil
}

func (userUsecase *UserUsecase) Delete(id string) error {
	err := userUsecase.UserRepository.Delete(id)

	if err != nil {
		log.Fatalf("Erro ao deletar o usuario: %v", err)
		return err
	}

	return nil
}
