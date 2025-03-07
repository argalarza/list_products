package main

import (
	"list-products/controllers"
	"list-products/models"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar la base de datos
	db, err := models.InitDB()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	// Crear un enrutador de mux
	r := mux.NewRouter()

	// Registrar las rutas para obtener productos
	r.HandleFunc("/products", controllers.GetAllProductsHandler(db)).Methods("GET")
	// Aplicar CORS
	handler := cors.Default().Handler(r)

	// Iniciar el servidor
	log.Println("Servidor iniciado en :3010")
	log.Fatal(http.ListenAndServe(":3010", handler))
}
