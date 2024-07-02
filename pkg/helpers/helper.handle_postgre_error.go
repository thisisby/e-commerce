package helpers

import (
	"database/sql"
	"errors"
	"ga_marketplace/internal/constants"
	"github.com/lib/pq"
)

func PostgresErrorTransform(err error) error {
	if err == nil {
		return nil
	}

	// duplicate key value violates unique constraint
	var pgErr *pq.Error
	ok := errors.As(err, &pgErr)
	if ok {
		switch pgErr.Code {
		case "23505":
			return constants.ErrRowExists
		case "23503":
			return constants.ErrForeignKeyViolation
		}
	}

	// no rows in result set
	if errors.Is(err, sql.ErrNoRows) {
		return constants.ErrRowNotFound
	}

	return err
}
