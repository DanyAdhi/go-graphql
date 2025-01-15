package pets

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() (*[]Pets, error)
	GetOne(id int32) (*Pets, error)
	Create(input NewPet) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() (*[]Pets, error) {
	var pets []Pets
	err := r.db.Table("pets p").
		Select(`p.id, p.name, p.type, p.ownerId, o.name as ownerName`).
		Joins(`LEFT JOIN owner o ON o.id = p.ownerId`).
		Scan(&pets).Error

	if err != nil {
		return nil, err
	}

	return &pets, nil
}

func (r repository) GetOne(id int32) (*Pets, error) {

	var pet Pets

	err := r.db.Table(`pets p`).
		Select(`p.id, p.name, p.type, p.ownerId, o.name as ownerName`).
		Joins(`LEFT JOIN owner o ON o.id = p.ownerId`).
		Where(`p.id = ?`, id).
		First(&pet).Error
	if err != nil {
		return nil, err
	}

	return &pet, nil
}

func (r repository) Create(input NewPet) (int, error) {
	var id int

	err := r.db.Raw(
		`INSERT INTO pets (name, type, "ownerId") VALUES ($1, $2, $3) RETURNING id`,
		input.Name,
		input.Type,
		input.OwnerId,
	).Scan(&id).Error
	if err != nil {
		return 0, err
	}

	return id, nil
}
