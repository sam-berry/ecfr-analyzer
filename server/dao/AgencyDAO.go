package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-berry/ecfr-analyzer/server/data"
	"time"
)

type AgencyDAO struct {
	Db *sql.DB
}

func (d *AgencyDAO) Create(
	ctx context.Context,
	agency *data.Agency,
) (*data.Agency, error) {
	id := uuid.New().String()

	_, err := d.Db.ExecContext(
		ctx,
		`INSERT INTO agency(agencyId, name, shortName, sortableName, slug, createdTimestamp) 
         VALUES ($1, $2, $3, $4, $5, $6)
         RETURNING id`,
		id,
		agency.Name,
		agency.ShortName,
		agency.SortableName,
		agency.Slug,
		time.Now().UTC(),
	)

	if err != nil {
		return nil, fmt.Errorf("error inserting agency: %v", err)
	}

	agency.Id = id
	return agency, nil
}

func (d *AgencyDAO) FindBySlug(
	ctx context.Context,
	slug string,
) (*data.Agency, error) {
	var a data.Agency

	err := d.Db.QueryRowContext(
		ctx,
		`SELECT id, agencyId, name, shortName, sortableName, slug
         FROM agency
         WHERE slug = $1`,
		slug,
	).Scan(&a.InternalId, &a.Id, &a.Name, &a.ShortName, &a.SortableName, &a.Slug)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error finding agency by slug: %v, %v", slug, err)
	}

	return &a, nil
}
