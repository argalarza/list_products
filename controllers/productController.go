package controllers

import (
	"encoding/json"
	"list-products/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// GetAllProductsHandler maneja la obtención de todos los productos
func GetAllProductsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := models.GetAllProducts(db)
		if err != nil {
			http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

// DetailProductHandler maneja la obtención de detalles de un producto
func DetailProductHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		product, err := models.GetProductByID(db, id)
		if err != nil {
			http.Error(w, "Error al obtener el producto", http.StatusInternalServerError)
			return
		}

		if product == nil {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}
