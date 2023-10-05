package main

import (
	"net/http"

	"github.com/Rethory/go-gorm-rest-api/db"
	"github.com/Rethory/go-gorm-rest-api/db/models"
	"github.com/Rethory/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
)



func main() {

	db.DBConexion()

	db.DB.AutoMigrate(models.Mascota{})
	db.DB.AutoMigrate(models.Propietario{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/propietarios", routes.GetPropietariosHandler).Methods("GET")
	r.HandleFunc("/propietarios", routes.PostPropietarioHandler).Methods("POST")
	r.HandleFunc("/propietarios/{id}", routes.GetPropietarioHandler).Methods("GET")
	r.HandleFunc("/propietarios/{id}", routes.DeletePropietarioHandler).Methods("DELETE")

	// Mascotas routes
	r.HandleFunc("/mascotas", routes.GetMascotasHandler).Methods("GET")
	r.HandleFunc("/mascotas", routes.AddMascotasHandler).Methods("POST")
	r.HandleFunc("/mascotas/{id}", routes.GetMascotaHandler).Methods("GET")
	r.HandleFunc("/mascotas/{id}", routes.DeleteMascotasHandler).Methods("DELETE")


	http.ListenAndServe(":3000", r)
}