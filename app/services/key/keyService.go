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
	Keys         map[string]int
}

func New(config config.Config) ApiKeychain {
	return &ApiKeyIml{
		DbConnection: config.Settings.KeyStoreDbConnection,
		Keys:         make(map[string]int),
	}
}

func (iml *ApiKeyIml) IsValid(apiKey *string) bool {
	ctx := context.Background()
	if *apiKey == "" {
		return false
	}
	_, exist := iml.Keys[*apiKey]
	if exist {
		return exist
	}
	log.Debugln("cache miss - fetching from db")
	con, err := sql.Open("postgres", iml.DbConnection)
	if err != nil || con.Ping() != nil {
		log.Fatalln("failed connecting to db")
	}
	q := db.New(con)
	dbRecord, err := q.GetApiKey(ctx, *apiKey)
	if err != nil {
		log.Errorln("failed get key", *apiKey)
	}
	iml.Keys[dbRecord.ApiKey] = 0
	return dbRecord.Enabled
}
