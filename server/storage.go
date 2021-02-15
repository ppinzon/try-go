package server

type User struct {
	ID int // Think about how server will generate unique IDs
	Name string
	Surname string
	Age int
}

type Storage interface {
	GetUser(id int) (*User, error)
	CreateUser(*User) error // Make sure ID is added to the struct OR return ID from this method
	DeleteUser(id int) error
}

func NewInMemoryStorage() Storage {
	// TODO: IMPLEMENT PART A
	var out Storage
	return out
}


