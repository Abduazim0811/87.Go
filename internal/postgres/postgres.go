package postgres

import (
	"database/sql"
	"fmt"

	"87.GO/internal/models"
	"github.com/Masterminds/squirrel"
)

type Item struct {
	db *sql.DB
}

func NewItem(db *sql.DB) *Item {
	return &Item{db: db}
}

func (i *Item) StoreNewItem(item *models.Item) (*models.Item, error) {
	var items models.Item
	sql, args, err := squirrel.
		Insert("items").
		Columns("name", "value").
		Values(item.Name, item.Value).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to insert item: %w", err)
	}

	row := i.db.QueryRow(sql, args...)
	if err := row.Scan(&items.ID); err != nil {
		return nil, fmt.Errorf("unable to scan item ID: %w", err)
	}

	return &items, nil
}

func (i *Item) StoreGetItem(id int) (*models.Item, error) {
	var item models.Item

	sql, args, err := squirrel.
		Select("*").From("items").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to get item: %w", err)
	}

	row := i.db.QueryRow(sql, args...)
	if err := row.Scan(&item.ID, &item.Name, &item.Value); err != nil {
		return nil, fmt.Errorf("unable to scan item: %w", err)
	}

	return &item, nil
}

func (i *Item) StoreUpdateItem(item *models.Item) (*models.Item, error) {
	sql, args, err := squirrel.
		Update("items").
		Set("name", item.Name).
		Set("value", item.Value).
		Where(squirrel.Eq{"id": item.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to update item: %w", err)
	}

	_, err = i.db.Exec(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to execute update: %w", err)
	}

	return item, nil
}

func (i *Item) StoreDeleteItem(id int) error {
	sql, args, err := squirrel.
		Delete("items").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("unable to delete item: %w", err)
	}

	_, err = i.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("unable to execute delete: %w", err)
	}

	return nil
}
