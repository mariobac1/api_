package person

import "github.com/mariobac1/api_/models"

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(p *models.Person) error {
	return uc.storage.Create(p)
}

func (uc UseCase) GetAll() (models.Persons, error) {
	return uc.storage.GetAll()
}

// func (uc UseCase) GetByID(ID uint) (models.Person, error) {
// 	return uc.storage.GetByID(ID)
// }
