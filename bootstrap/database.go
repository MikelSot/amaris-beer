package bootstrap

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"

	"github.com/MikelSot/amaris-beer/model"
)

func newDatabase(ctx context.Context, dbConfig model.DatabaseConfig, applicationName string) model.PgxPool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("could not parse config pgxpool, err: %v", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = applicationName

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("could not connection to db, err: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("could ping database, err: %v", err)
	}

	return db
}
