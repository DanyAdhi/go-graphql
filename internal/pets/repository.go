package pets

import (
	"database/sql"
	"log"
)

type Repository interface {
	GetAll() (*[]Pets, error)
	GetOne(id int32) (*Pets, error)
	Create(input NewPet) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() (*[]Pets, error) {
	row, err := r.db.Query(
		`select p.id, p."name", p.type, p."ownerId", o."name" as "ownerName"  from pet p 
			left join "owner" o on o.id = p."ownerId" `,
	)
	if err != nil {
		log.Printf("Error get pets. %v", err)
		return nil, err
	}
	defer row.Close()

	var pets []Pets
	for row.Next() {
		var data Pets

		err := row.Scan(&data.ID, &data.Name, &data.Type, &data.OwnerId, &data.OwnerName)
		if err != nil {
			log.Printf("Row scan error: %v", err)
			return nil, err
		}

		pets = append(pets, data)
	}

	return &pets, nil
}

func (r repository) GetOne(id int32) (*Pets, error) {
	row := r.db.QueryRow(
		`SELECT p.id, p."name", p.type, p."ownerId", o."name" as "ownerName"  FROM pet p 
			LEFT JOIN "owner" o on o.id = p."ownerId" WHERE p.id = $1`,
		id,
	)

	var data Pets
	err := row.Scan(&data.ID, &data.Name, &data.Type, &data.OwnerId, &data.OwnerName)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r repository) Create(input NewPet) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO pet (name, type, "ownerId") VALUES ($1, $2, $3) RETURNING id`,
		input.Name,
		input.Type,
		input.OwnerId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
