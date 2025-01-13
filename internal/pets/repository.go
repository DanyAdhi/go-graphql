package pets

import (
	"database/sql"
	"log"
)

type Repository interface {
	GetAll() (*[]Pets, error)
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
		log.Printf("Error get total record articles. %v", err)
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