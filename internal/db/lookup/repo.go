package lookup

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/prybintsev/stakefish/internal/models"
)

type Repo struct {
	db *sql.DB
}

func NewLookupRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) InsertLookup(ctx context.Context, lookup models.Lookup) error {
	resp, err := json.Marshal(lookup)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, "INSERT INTO lookup_history (domain, response, created_at) VALUES ($1, $2, $3)", lookup.Domain, string(resp), lookup.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetLastLookups(ctx context.Context) ([]models.Lookup, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT response FROM lookup_history ORDER BY created_at DESC LIMIT 20")
	if err != nil {
		return nil, err
	}

	lookups := make([]models.Lookup, 0)
	for rows.Next() {
		var res []byte
		var lookup models.Lookup
		err = rows.Scan(&res)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(res, &lookup)
		if err != nil {
			return nil, err
		}
		lookups = append(lookups, lookup)
	}

	return lookups, nil
}
