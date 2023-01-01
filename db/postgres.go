package db
import (
	"github.com/jackc/pgx"
)
func InitDB(config) PostgresClient {
	client := Postgres({
		host=config.host,
		user=config.user,
		port=config.port,
		db=config.db
	})
	return client
}