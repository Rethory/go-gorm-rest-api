package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Rethory/go-gorm-rest-api/db"
	"github.com/Rethory/go-gorm-rest-api/db/models"
	"github.com/gorilla/mux"
)

func GetMascotasHandler(w http.ResponseWriter, r *http.Request) {
	var mascotas []models.Mascota
	db.DB.Find(&mascotas)
	json.NewEncoder(w).Encode(mascotas)
}

func AddMascotasHandler(w http.ResponseWriter, r *http.Request) {
	var mascota models.Mascota
	json.NewDecoder(r.Body).Decode(&mascota)
	mascotaCreada := db.DB.Create(&mascota)
	err := mascotaCreada.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&mascota)

}

func GetMascotaHandler(w http.ResponseWriter, r *http.Request) {
	var mascota models.Mascota
	params := mux.Vars(r)

	db.DB.First(&mascota, params["id"])

	if mascota.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Mascota no encontrada"))
		return
	}

	json.NewEncoder(w).Encode(&mascota)
}

func DeleteMascotasHandler(w http.ResponseWriter, r *http.Request) {
	var mascota models.Mascota
	params := mux.Vars(r)

	db.DB.First(&mascota, params["id"])

	if mascota.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Mascota no encontrada"))
		return
	}

	db.DB.Unscoped().Delete(&mascota)
	w.WriteHeader(http.StatusNoContent) // 204
}