package pets

import (
	"errors"
	"fmt"
)

type Service interface {
	GetAll() (*[]Pets, error)
	GetOne(id int32) (*Pets, error)
	Create(input NewPet) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetAll() (*[]Pets, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s service) GetOne(id int32) (*Pets, error) {
	pet, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

// func (s service) Create(input NewPet) error {
// 	err := s.repo.Create(input)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s service) Create(data pets.NewPet) (*pets.Pet, error) {
// 	newPet := pets.Pet{
// 		Name:    data.Name,
// 		Type:    data.Type,
// 		OwnerId: data.OwnerId,
// 	}
// 	err := s.repo.Create(&newPet).Error
// 	if err != nil {
// 		return nil, nil
// 	}
// 	return &newPet, nil
// }

func (s service) Create(input NewPet) (int, error) {
	// Validasi input
	if input.Name == "" || input.Type == "" || input.OwnerId <= 0 {
		return 0, errors.New("invalid input data")
	}

	// Panggil repository untuk menyimpan data
	id, err := s.repo.Create(input)
	if err != nil {
		return 0, fmt.Errorf("failed to create pet: %w", err)
	}

	return id, nil
}
