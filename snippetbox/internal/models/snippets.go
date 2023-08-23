package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID 			int
	Title		string
	Content	string
	Created	time.Time
	Expires	time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, CONVERT_TZ(UTC_TIMESTAMP(), '+00:00', '+02:00'), DATE_ADD(CONVERT_TZ(UTC_TIMESTAMP(), '+00:00', '+02:00'), INTERVAL ? DAY))`

	result, err := model.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return int(id), nil
}

func (model *SnippetModel) Get(id int) (*Snippet, error) {
	
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CONVERT_TZ(UTC_TIMESTAMP(), '+00:00', '+02:00') AND id = ?`
	
	row := model.DB.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (model *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}