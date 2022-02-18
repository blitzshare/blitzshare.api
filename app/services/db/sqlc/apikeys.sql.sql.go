// Code generated by sqlc. DO NOT EDIT.
// source: apikeys.sql.sql

package db

import (
	"context"
)

const getApiKey = `-- name: GetApiKey :one
SELECT id, created_at, api_key, enabled FROM "api_keys" WHERE api_key = $1
`

func (q *Queries) GetApiKey(ctx context.Context, apiKey string) (ApiKey, error) {
	row := q.db.QueryRowContext(ctx, getApiKey, apiKey)
	var i ApiKey
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ApiKey,
		&i.Enabled,
	)
	return i, err
}
