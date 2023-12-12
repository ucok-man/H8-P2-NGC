package entity

import "database/sql"

type Entity struct {
	Inventory InventoryEntity
}

func New(db *sql.DB) *Entity {
	return &Entity{
		Inventory: InventoryEntity{db: db},
	}
}
