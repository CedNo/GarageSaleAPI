package services

import (
	"GarageSaleAPI/domain/user"
	"GarageSaleAPI/infrastructure/persistence/memory"
	"GarageSaleAPI/interfaces/dto"
	"errors"
	"log/slog"
	"net/http"
	"net/mail"
	"time"
)

func AddUser(userDTO dto.UserDTO) (error, int) {
	email, err := mail.ParseAddress(userDTO.Email)
	if err != nil {
		slog.Error("error parsing email")
		return errors.New("bad request email"), http.StatusBadRequest
	}

	newUser := user.CreateUser(userDTO.Id, userDTO.Username, userDTO.Password, *email, time.Now())
	memory.AddUser(newUser)

	return nil, http.StatusCreated
}

func GetUserByUsername(username string) (*user.User, error) {
	u, err := memory.GetUserById(username)
	if err != nil {
		return nil, err
	}

	return u, nil
}
