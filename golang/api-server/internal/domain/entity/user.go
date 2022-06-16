package entity

import "time"

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string) *User {
	timeNow := time.Now()
	return &User{
		Name:      name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

func (u *User) UpdateUser(name string) *User {
	timeNow := time.Now()
	u.Name = name
	u.UpdatedAt = timeNow
	return u
}
