package database

import (
	"database/sql"
	"errors"

	"github.com/zHenriqueGN/CourseCLI/internal/entity"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) Create(category *entity.Category) error {
	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.ID, category.Name, category.Description)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	stmt, err := c.db.Prepare("SELECT id, name, description FROM categories WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category entity.Category
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}
	return nil, ErrCategoryNotFound
}

func (c *CategoryDB) FindAll() ([]*entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
