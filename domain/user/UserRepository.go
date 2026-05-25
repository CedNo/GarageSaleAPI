package user

type UserRepository interface {
	add(user User) (User, error)
}
