package entity

import (
	"database/sql"
)

type Inventory struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ItemCode    string `json:"item_code"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type InventoryEntity struct {
	db *sql.DB
}

func (e *InventoryEntity) GetAll() ([]*Inventory, error) {
	query := `
		SELECT id, name, item_code, stock, description, status
		FROM inventories
	`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []*Inventory
	for rows.Next() {
		var inventory = Inventory{}

		err := rows.Scan(
			&inventory.ID,
			&inventory.Name,
			&inventory.ItemCode,
			&inventory.Stock,
			&inventory.Description,
			&inventory.Status,
		)
		if err != nil {
			return nil, err
		}

		inventories = append(inventories, &inventory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventories, nil
}

func (e *InventoryEntity) GetByID(id int64) (*Inventory, error) {
	query := `
		SELECT id, name, item_code, stock, description, status
		FROM inventories
		WHERE id = ?
	`

	var inventory Inventory
	err := e.db.QueryRow(query, id).Scan(
		&inventory.ID,
		&inventory.Name,
		&inventory.ItemCode,
		&inventory.Stock,
		&inventory.Description,
		&inventory.Status,
	)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (e *InventoryEntity) Insert(inventory *Inventory) (int64, error) {
	query := `
		INSERT INTO inventories(name, item_code, stock, description, status)
		VALUES (?, ?, ?, ?, ?)
	`
	args := []any{inventory.Name, inventory.ItemCode, inventory.Stock, inventory.Description, inventory.Status}

	result, err := e.db.Exec(query, args...)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (e *InventoryEntity) UpdateByID(inventory *Inventory) error {
	query := `
		UPDATE inventories
		SET
			name = ?,
			item_code = ?,
			stock = ?,
			description = ?,
			status = ?
		WHERE 
			id = ?
	`

	args := []any{inventory.Name, inventory.ItemCode, inventory.Stock, inventory.Description, inventory.Status, inventory.ID}

	_, err := e.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (e *InventoryEntity) DeleteByID(id int64) error {
	query := `
		DELETE FROM inventories
		WHERE id = ?
	`

	_, err := e.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
