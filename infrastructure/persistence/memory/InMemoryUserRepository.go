package memory

import (
	"GarageSaleAPI/domain/user"
	"errors"
)

var userList []user.User

func AddUser(user user.User) {
	userList = append(userList, user)
}

func GetUserById(username string) (*user.User, error) {
	for _, foundUser := range userList {
		if foundUser.Username == username {
			return &foundUser, nil
		}
	}
	return nil, errors.New("User not found")
}
