package dependencyinversion

// Dependency inversion says that a higher level class should depend on a lower level class
// but the abstraction between them
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
