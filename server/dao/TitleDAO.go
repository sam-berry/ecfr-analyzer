package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"strings"
	"time"
)

var InvalidXMLErrorCode = pq.ErrorCode("2200N")

type TitleDAO struct {
	Db *sql.DB
}

func (d *TitleDAO) CountWords(ctx context.Context, agencyName string, titles []int) (int, error) {
	var count int
	err := d.Db.QueryRowContext(
		ctx,
		`SELECT 
          SUM(
            COALESCE(
              ARRAY_LENGTH(
                ARRAY_REMOVE(
                  REGEXP_SPLIT_TO_ARRAY(
                    TRIM((XPATH(
                      'string(//HEAD[contains(translate(., "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"), "' || $1 || '")]/..)',
                      content
                    ))[1]::TEXT),
                    '\s+'
                  ),
                  ''
                ),
                1
              ),
              0
            )
       ) AS total_word_count FROM title WHERE name = ANY($2);`,
		strings.ToLower(agencyName),
		pq.Array(titles),
	).Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("error counting words for agency, %v, %w", agencyName, err)
	}

	return count, nil
}

func (d *TitleDAO) Insert(
	ctx context.Context,
	name int,
	content []byte,
) error {
	id := uuid.New().String()

	err := d.insertTitle(ctx, name, content, id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == InvalidXMLErrorCode {
				log.Info(
					fmt.Sprintf(
						"Invalid XML detected, attempting to scrub title %v",
						name,
					),
				)
				err = d.insertTitle(ctx, name, scrubXML(content), id)
				if err == nil {
					return nil
				}

				var pqErr *pq.Error
				if errors.As(err, &pqErr) {
					if pqErr.Code == InvalidXMLErrorCode {
						log.Info(
							fmt.Sprintf(
								"Invalid XML detected, attempting to aggressively scrub title %v",
								name,
							),
						)
						scrubbedXML, err := scrubXMLAggresive(content)
						if err != nil {
							return fmt.Errorf(
								"error aggressively scrubbing title, %v, %w",
								name,
								err,
							)
						}
						err = d.insertTitle(ctx, name, scrubbedXML, id)
						if err == nil {
							return nil
						}

						return fmt.Errorf(
							"error inserting title after aggresive scrub, %v, %w",
							name,
							err,
						)
					} else {
						return fmt.Errorf("error inserting title after scrub, %v, %w", name, err)
					}
				} else {
					return fmt.Errorf("error inserting title after scrub, %v, %w", name, err)
				}
			}
		}
		return fmt.Errorf("error inserting title, %v, %w", name, err)
	}

	return nil
}

func (d *TitleDAO) insertTitle(ctx context.Context, name int, content []byte, id string) error {
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
