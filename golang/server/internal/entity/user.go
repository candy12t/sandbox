package entity

type User struct {
	ID   int
	Name string
}

var id int

func NewUser(name string) *User {
	id++
	return &User{
		ID:   id,
		Name: name,
	}
}
