package database

import (
	"database/sql"
	"fmt"

	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/repository"
)

type TagRepository struct {
	db *sql.DB
}

var _ repository.Tag = &TagRepository{}

func NewTagRepository(db *sql.DB) repository.Tag {
	return &TagRepository{
		db: db,
	}
}

func (tr *TagRepository) Save(tag *entity.Tag) (int, error) {
	result, err := tr.db.Exec("INSERT INTO tags (name) VALUES (?)", tag.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (tr *TagRepository) FindById(id int) (*entity.Tag, error) {
	var tag entity.Tag
	row := tr.db.QueryRow("SELECT * FROM tags WHERE id = ? AND delete_mark = 0", id)
	if err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeleteMark); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Not found tag by %v", id)
		}
		return nil, err
	}
	return &tag, nil
}

func (tr *TagRepository) FindAll() ([]*entity.Tag, error) {
	var tags []*entity.Tag
	rows, err := tr.db.Query("SELECT * FROM tags WHERE delete_mark = 0")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag entity.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeleteMark); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tags, nil
}

func (tr *TagRepository) Update(tag *entity.Tag) (int, error) {
	_, err := tr.db.Exec("UPDATE tags SET name = ? WHERE id = ? AND delete_mark = 0", tag.Name, tag.ID)
	if err != nil {
		return 0, err
	}
	return tag.ID, nil
}

func (tr *TagRepository) Delete(tag *entity.Tag) error {
	_, err := tr.db.Exec("UPDATE tags SET delete_mark = ? WHERE id = ?", tag.DeleteMark, tag.ID)
	if err != nil {
		return err
	}
	return nil
}
