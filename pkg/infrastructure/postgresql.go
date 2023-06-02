package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq" // pgx also supported
	"os"
)

type ConnDB struct {
	pool *pgxpool.Pool
}

func GenerateDSN(driver string) string {
	dsn := driver + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	return dsn
}

func OpenDB() (*pgxpool.Pool, error) {
	dsn := GenerateDSN("postgres://")
	// this returns connection pool
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	return dbPool, nil
}

func (p *ConnDB) Close() {
	p.pool.Close()
}
