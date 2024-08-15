package handler

import (
	"Auth-Server/internal/data/database"
	"Auth-Server/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type regParams struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"address"`
}

type LoginParams struct {
	LoginIdentifier string `json:"loginIdentifier"`
	Password        string `json:"password"`
}

func (s *DbConfig) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var regParams regParams

	if err := json.NewDecoder(r.Body).Decode(&regParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regParams.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user, err := s.q.CreateUserWithProfile(r.Context(), database.CreateUserWithProfileParams{
		ID:           uuid.New(),
		FirstName:    sql.NullString{String: regParams.FirstName, Valid: true},
		LastName:     sql.NullString{String: regParams.LastName, Valid: true},
		Username:     regParams.Username,
		Email:        regParams.Email,
		PasswordHash: string(hashedPassword),
		PhoneNumber:  sql.NullString{String: regParams.PhoneNumber, Valid: true},
		Address:      sql.NullString{String: regParams.Address, Valid: true},
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}
	utils.RespondWithJSON(w, http.StatusAccepted, user)
}

func (s *DbConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginParams LoginParams

	if err := json.NewDecoder(r.Body).Decode(&loginParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := s.q.GetUserByUsernameOrEmail(r.Context(), loginParams.LoginIdentifier)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginParams.Password))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	userProfile, err := s.q.GetUserByID(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error getting user")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, userProfile)
}
