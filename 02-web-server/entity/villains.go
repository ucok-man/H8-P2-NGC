package entity

import "database/sql"

type Villain struct {
	ID       int64
	Name     string
	Universe string
	ImageURL string
}

type VillainEntity struct {
	db *sql.DB
}

func (e *VillainEntity) GetAll() ([]*Villain, error) {
	query := `
		SELECT id, name, universe, image_url
		FROM villains
	`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}

	var villains []*Villain
	for rows.Next() {
		var villain = Villain{}

		err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			return nil, err
		}

		villains = append(villains, &villain)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return villains, nil
}
