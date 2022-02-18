package key

import (
	"blitzshare.api/app/config"
	db "blitzshare.api/app/services/db/sqlc"
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type ApiKeychain interface {
	IsValid(apiKey *string) bool
}

type ApiKeyIml struct {
	DbConnection string
}

func New(config config.Config) ApiKeychain {
	return &ApiKeyIml{
		DbConnection: config.Settings.KeyStoreDbConnection,
	}
}

func (iml *ApiKeyIml) IsValid(apiKey *string) bool {
	ctx := context.Background()
	result, err := sql.Open("postgres", iml.DbConnection)
	if err != nil {
		log.Fatalln("failed connecting to db")
	}
	q := db.New(result)
	dbApiKey, err := q.GetApiKey(ctx, *apiKey)
	return dbApiKey.Enabled
}
