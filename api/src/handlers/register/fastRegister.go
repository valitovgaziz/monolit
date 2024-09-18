package register

import (
	"api/src/models"
	"api/src/storage"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func FastRegister(w http.ResponseWriter, r *http.Request) {
	var ShortCredentials models.ShortCredentials

	// Decode body
	err := json.NewDecoder(r.Body).Decode(&ShortCredentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(ShortCredentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := uuid.New()

	user := models.User{
		Id:       id,
		Login:    ShortCredentials.Phone,
		Email:    "",
		Password: hashedPassword,
		Phone:    ShortCredentials.Phone,
		Role:     ShortCredentials.Role,
	}
	result := storage.PSQL_GORM_DB.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
