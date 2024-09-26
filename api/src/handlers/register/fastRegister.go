package register

import (
	"api/src/models"
	"api/src/storage"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/profclems/go-dotenv"
)

// Fast regisgter handlers by Phone and password only.
func FastRegister(w http.ResponseWriter, r *http.Request) {
	var Phone models.Phone

	// Decode body
	err := json.NewDecoder(r.Body).Decode(&Phone)
	if err != nil {
		slog.Info("Struct from request not decodet", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := hashPassword(Phone.Password)
	if err != nil {
		slog.Info("Can't hash password", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create new UUID
	id := uuid.New()

	// create account
	account := models.Account{
		Id:       id,
		Login:    Phone.Phone,
		Password: hashedPassword,
		Phone:    Phone.Phone,
		Role:     dotenv.GetString("ROLE"),
	}

	// save to DB
	result := storage.PSQL_GORM_DB.Create(&account)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// log
	slog.Info("Account saved")

	// response status write OK
	w.WriteHeader(http.StatusOK)

}
