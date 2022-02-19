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
	log.Debugln("Connecting to database", iml.DbConnection)
	con, err := sql.Open("postgres", iml.DbConnection)
	if err != nil || con.Ping() != nil {
		log.Fatalln("failed connecting to db")
	}
	q := db.New(con)
	allKeys, err := q.GetApiKeys(ctx)
	log.Debugln("allKeys", len(allKeys), allKeys, err)
	dbApiKey, err := q.GetApiKey(ctx, *apiKey)
	if err != nil {
		log.Fatalln("failed get key", *apiKey)
	}
	log.Debugln("dbApiKey", dbApiKey)

	return dbApiKey.Enabled
}
