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

func NewUserRepository(db *sql.DB) repository.User {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(user *entity.User) (*entity.User, error) {
	result, err := ur.db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return ur.FindById(int(id))
}

func (ur *UserRepository) FindById(id int) (*entity.User, error) {
	var user entity.User
	err := ur.db.QueryRow("SELECT * FROM users WHERE id = ? AND delete_mark = 0", id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeleteMark)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Not found user by %v", id)
	}
	return &user, nil
}

func (ur *UserRepository) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	rows, err := ur.db.Query("SELECT * FROM users WHERE delete_mark = 0")
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
	_, err := ur.db.Exec("UPDATE users SET name = ? WHERE id = ? AND delete_mark = 0", user.Name, user.ID)
	if err != nil {
		return nil, err
	}
	return ur.FindById(user.ID)
}

func (ur *UserRepository) Delete(user *entity.User) error {
	_, err := ur.db.Exec("UPDATE users SET delete_mark = ? WHERE id = ?", user.DeleteMark, user.ID)
	if err != nil {
		return err
	}
	return nil
}
