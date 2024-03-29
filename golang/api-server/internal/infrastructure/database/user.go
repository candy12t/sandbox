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

func (ur *UserRepository) Save(user *entity.User) (int, error) {
	result, err := ur.db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (ur *UserRepository) FindById(id int) (*entity.User, error) {
	var user entity.User
	row := ur.db.QueryRow("SELECT * FROM users WHERE id = ? AND delete_mark = 0", id)
	if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeleteMark); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Not found user by %v", id)
		}
		return nil, err
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

func (ur *UserRepository) Update(user *entity.User) (int, error) {
	_, err := ur.db.Exec("UPDATE users SET name = ? WHERE id = ? AND delete_mark = 0", user.Name, user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (ur *UserRepository) Delete(user *entity.User) error {
	_, err := ur.db.Exec("UPDATE users SET delete_mark = ? WHERE id = ?", user.DeleteMark, user.ID)
	if err != nil {
		return err
	}
	return nil
}
