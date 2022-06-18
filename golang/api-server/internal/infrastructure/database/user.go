package database

import (
	"database/sql"
	"fmt"

	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/repository"
)

type UserRepository struct {
	db *sql.DB
}

var _ repository.User = &UserRepository{}

var cacheUsers []*entity.User

var id int

func NewUserRepository(db *sql.DB) repository.User {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(user *entity.User) (*entity.User, error) {
	id++
	user.ID = id
	cacheUsers = append(cacheUsers, user)
	return user, nil
}

func (ur *UserRepository) FindById(id int) (*entity.User, error) {
	var user entity.User
	err := ur.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeleteMark)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Not found user by %v", id)
	}
	return &user, nil
}

func (ur *UserRepository) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	rows, err := ur.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeleteMark); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Update(user *entity.User) (*entity.User, error) {
	for _, u := range cacheUsers {
		if u.ID == user.ID && !user.DeleteMark {
			u = user
		}
	}
	return user, nil
}

func (ur *UserRepository) Delete(user *entity.User) error {
	for _, u := range cacheUsers {
		if u.ID == user.ID {
			u = user
		}
	}
	return nil
}
