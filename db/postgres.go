package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/preemptcordon/server/obj"
)

var DBPool pgxpool.Pool

func InitDB(config obj.ServerSettings) {
	dbpool, err := pgxpool.New(context.Background(), config.PostgresURI)
	if err != nil {
		panic(err)
	}
	DBPool = *dbpool
	fmt.Println("DB initialized")
}
