package dependencyinversion

// DIP states that high-level modules should not depend on low-level modules; both should depend on abstractions.
// Abstractions should not depend on details; details should depend on abstractions.

type User struct {
	Username string
}

type IReadUser interface {
	GetUser() User
}

type UserReadStore struct{}

func NewUserRead() IReadUser {
	return &UserReadStore{}
}

func (*UserReadStore) GetUser() User {
	return User{
		Username: "pyro",
	}
}

func (*UserReadStore) CreateUser() User {
	return User{
		Username: "pyro",
	}
}

func UserHandler() {

	userStore := NewUserRead()

	userStore.GetUser()

	// this function doesn't work because there is no abstraction implementation of it
	// userStore.CreateUser()
}
