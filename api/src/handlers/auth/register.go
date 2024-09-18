package auth

import (
	"api/src/models"
	"api/src/storage"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var Credentials models.Credentials
	// Decode body
	err := json.NewDecoder(r.Body).Decode(&Credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// hash password
	hashedPassword, err := hashPassword(Credentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := uuid.New()

	user := models.User{
		Id:       id,
		Login:    Credentials.Login,
		Email:    Credentials.Email,
		Password: hashedPassword,
		Phone:    Credentials.Phone,
		Role:     Credentials.Role,
	}
	result := storage.PSQL_GORM_DB.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
