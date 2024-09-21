package register

import (
	"api/src/models"
	"api/src/storage"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func FastRegister(w http.ResponseWriter, r *http.Request) {
	slog.Info("FastRegister")
	var ShortCredentials models.ShortCredentials

	// Decode body
	err := json.NewDecoder(r.Body).Decode(&ShortCredentials)
	if err != nil {
		slog.Info("Not decoded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(ShortCredentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := uuid.New()

	account := models.Account{
		Id:       id,
		Login:    ShortCredentials.Phone,
		Password: hashedPassword,
		Phone:    ShortCredentials.Phone,
		Role:     ShortCredentials.Role,
	}
	result := storage.PSQL_GORM_DB.Create(&account)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
