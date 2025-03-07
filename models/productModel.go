package models

import (
	"github.com/jmoiron/sqlx"
)

// Product representa la estructura de un producto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
}

// GetProductByID obtiene los detalles de un producto por su ID
func GetProductByID(db *sqlx.DB, id int) (*Product, error) {
	var product Product
	query := `SELECT id, name, description, price, quantity, category, brand FROM Products WHERE id = ?`
	err := db.Get(&product, query, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProducts obtiene todos los productos
func GetAllProducts(db *sqlx.DB) ([]Product, error) {
	var products []Product
	query := `SELECT id, name, description, price, quantity, category, brand FROM Products`
	err := db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
