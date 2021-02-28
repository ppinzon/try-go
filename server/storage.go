package server

//UserMap bla
// var UserMap = make(map[int]*User)
// var NextUserID int = 1

// User bla
type (
	User struct {
		ID      int    `json:"id" xml:"id" form:"id" query:"id"` // Think about how server will generate unique IDs
		Name    string `json:"name" xml:"name" form:"name" query:"name"`
		Surname string `json:"surname" xml:"surname" form:"surname" query:"surname"`
		Age     int    `json:"age" xml:"age" form:"age" query:"age"`
	}
	handler struct {
		db map[int]*User
	}
)

// Storage bla
type Storage interface {
	GetUser(id int) (*User, error)
	CreateUser(*User) error // Make sure ID is added to the struct OR return ID from this method
	DeleteUser(id int) error
}

// NewInMemoryStorage bla
func NewInMemoryStorage() Storage {
	// TODO: IMPLEMENT PART A
	var out Storage
	return out
}
