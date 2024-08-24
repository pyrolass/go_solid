package interfacesegregation

// Interface segregation says that we should separate our interfaces
// for example we should have different interfaces for read and write methods

type User struct {
	Username string
}

type IReadUser interface {
	GetUser() User
}

type IWriteUser interface {
	CreateUser()
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

type UserWriteStore struct{}

func NewUserWrite() IWriteUser {
	return &UserWriteStore{}
}

func (*UserWriteStore) CreateUser() {

}
