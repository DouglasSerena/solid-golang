package main

import "log"

type IAnimalRepository interface {
	Save(anima Animal)
}

type AnimalRepositoryMemory struct {
	IAnimalRepository
}

func (repo AnimalRepositoryMemory) Save(animal Animal) {
	log.Printf("Animal %s Salvo", animal.Name)
}

type IAnimal interface {
	Mover()
	Comer()
	Save()
}
type Animal struct {
	IAnimal
	Repo IAnimalRepository
	Name string
}

func (animal *Animal) Save() {
	animal.Repo.Save(*animal)
}
func (animal *Animal) Mover() {
	log.Println("Movindo")
}
func (animal *Animal) Comer() {
	log.Println("Comendo algo")
}

type Dog struct {
	Animal
}

func (dog *Dog) Comer() {
	log.Println("Comendo ração")
}

func main() {
	repo := AnimalRepositoryMemory{}
	dog := Dog{Animal{Repo: repo, Name: "tobias"}}

	log.Println(dog.Name)
	dog.Mover()
	dog.Comer()
	dog.Save()
}
