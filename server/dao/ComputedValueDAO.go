package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-berry/ecfr-analyzer/server/data"
	"time"
)

type ComputedValueDAO struct {
	Db *sql.DB
}

func (d *ComputedValueDAO) Insert(
	ctx context.Context,
	cv *data.ComputedValue,
) error {
	id := uuid.New().String()

	dBytes, err := cv.Data.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error converting data to bytes: %v", err)
	}

	_, err = d.Db.ExecContext(
		ctx,
		`INSERT INTO computed_value(valueId, key, data, createdTimestamp) 
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (key) DO UPDATE
         SET data = $3, createdTimestamp = $4
         WHERE computed_value.key = $2`,
		id,
		cv.Key,
		dBytes,
		time.Now().UTC(),
	)

	if err != nil {
		return fmt.Errorf("error inserting computed value, %v, %w", cv.Key, err)
	}

	return nil
}
