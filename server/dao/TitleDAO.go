package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	err := d.insertTitle(ctx, name, content, id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "2200N" { // invalid XML error
				return d.insertTitle(ctx, name, scrubXML(content), id)
			}
		}
		return fmt.Errorf("error inserting title, %v, %w", name, err)
	}

	return nil
}

func (d *TitleDAO) insertTitle(ctx context.Context, name string, content []byte, id string) error {
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
	return err
}
