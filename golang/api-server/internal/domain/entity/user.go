package entity

import "time"

type User struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteMark bool
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

func (u *User) DeleteUser() *User {
	timeNow := time.Now()
	u.UpdatedAt = timeNow
	u.DeleteMark = true
	return u
}
