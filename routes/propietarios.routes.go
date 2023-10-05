package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Rethory/go-gorm-rest-api/db"
	"github.com/Rethory/go-gorm-rest-api/db/models"
	"github.com/gorilla/mux"
)

func GetPropietariosHandler(w http.ResponseWriter, r *http.Request) {
	var propietarios []models.Propietario
	db.DB.Find(&propietarios)
	json.NewEncoder(w).Encode(&propietarios)
}

func GetPropietarioHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario
	params := mux.Vars(r)
	db.DB.First(&propietario, params["id"])

	if propietario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Propietario no encontrado"))
		return
	}

	db.DB.Model(&propietario).Association("Mascotas").Find(&propietario.Mascotas)

	json.NewEncoder(w).Encode(&propietario)
}

func PostPropietarioHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario
	json.NewDecoder(r.Body).Decode(&propietario)
	propietarioCreado := db.DB.Create(&propietario)
	err := propietarioCreado.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&propietario)
}

func DeletePropietarioHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario
	params := mux.Vars(r)
	db.DB.First(&propietario, params["id"])

	if propietario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Propietario con encontrado"))
		return
	}

	db.DB.Unscoped().Delete(&propietario)
	w.WriteHeader(http.StatusOK)
}