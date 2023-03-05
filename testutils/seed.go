package testutils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func (ts *TestSuite) createTables() error {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, ts.Config.Database.DSN)
	if err != nil {
		return err
	}

	defer db.Close(ctx)

	_, err = db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS articles(
		id SERIAL PRIMARY KEY,
		uuid VARCHAR(50) NOT NULL,
		title VARCHAR(50) NOT NULL,
		content TEXT,
		author VARCHAR(50) NOT NULL,
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
		updated_at timestamp NULL DEFAULT NULL,
		deleted_at timestamp NULL DEFAULT NULL
	)`)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TestSuite) seedPostgres() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, ts.Config.Database.DSN)
	if err != nil {
		panic(err)
	}

	defer db.Close(ctx)

	// seed data
	seedDataRaw, err := os.ReadFile(ts.Seed)
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(ctx, string(seedDataRaw)); err != nil {
		panic(err)
	}
}

func (ts *TestSuite) clearTable() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, ts.Config.Database.DSN)
	if err != nil {
		panic(err)
	}

	defer db.Close(ctx)

	_, err = db.Exec(ctx, `DELETE FROM articles;`)
	if err != nil {
		panic(err)
	}
}
