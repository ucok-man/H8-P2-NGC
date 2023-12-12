package entity

import "database/sql"

type Hero struct {
	ID       int64
	Name     string
	Universe string
	Skill    string
	ImageURL string
}

type HeroEntity struct {
	db *sql.DB
}

func (e *HeroEntity) GetAll() ([]*Hero, error) {
	query := `
		SELECT id, name, universe, skill, image_url
		FROM heroes
	`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}

	var heroes []*Hero
	for rows.Next() {
		var hero = Hero{}

		err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			return nil, err
		}

		heroes = append(heroes, &hero)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return heroes, nil
}
