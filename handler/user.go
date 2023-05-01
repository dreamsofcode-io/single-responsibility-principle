package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dreamsofcode-io/srp/email"
	"github.com/dreamsofcode-io/srp/hasher"
	"github.com/dreamsofcode-io/srp/user"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	repo           *user.UserRepository
	passwordHasher *hasher.Password
	mailer         *email.SESMailer
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encodedPassword, err := h.passwordHasher.HashPassword(data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := user.User{
		ID:           uuid.New(),
		Email:        data.Email,
		PasswordHash: encodedPassword,
	}

	if err := h.repo.Save(r.Context(), u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.mailer.SendEmail(r.Context(),
    u.Email,
		"Welcome to the service!",
		"You have successfully signed up for the service.",
	)

	w.WriteHeader(http.StatusCreated)
}
