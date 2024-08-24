package interfacesegregation

// ISP advocates for many client-specific interfaces rather than one general-purpose interface.
// Go's lightweight interfaces are perfect for this.

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
