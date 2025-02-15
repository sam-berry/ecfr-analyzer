package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TitleDAO struct {
	Db *sql.DB
}

func (d *TitleDAO) Insert(
	ctx context.Context,
	name string,
	content []byte,
) error {
	id := uuid.New().String()

	_, err := d.Db.ExecContext(
		ctx,
		`INSERT INTO title(titleId, name, content, createdTimestamp) 
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (name) DO UPDATE
         SET content = $3, createdTimestamp = $4
         WHERE title.name = $2`,
		id,
		name,
		content,
		time.Now().UTC(),
	)

	if err != nil {
		return fmt.Errorf("error inserting title, %v, %w", name, err)
	}

	return nil
}
