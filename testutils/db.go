package testutils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type testDB struct {
	Postgres *DBContainer
}

// NewTestDB creates test db containers
func NewTestDB() *testDB {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	if err := pool.Client.Ping(); err != nil {
		panic(err)
	}

	postgresTest := newPostgresTest(pool)
	db := &testDB{
		Postgres: postgresTest,
	}

	return db
}

type DBContainer struct {
	DSN      string
	resource *dockertest.Resource
}

func newPostgresTest(pool *dockertest.Pool) *DBContainer {
	var (
		dbname   = "postgres"
		user     = "postgres"
		password = "postgres"
	)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.1-alpine",
		Env: []string{
			"POSTGRES_DB=" + dbname,
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.Memory = 256 * 1024 * 1024
		config.CPUPercent = 10
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start postgres resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, hostAndPort, dbname)

	err = resource.Expire(600)
	if err != nil {
		log.Fatalf("Could not set resource expire: %s", err)
	}

	pool.MaxWait = 600 * time.Second
	if err = pool.Retry(func() error {
		ctx := context.Background()
		db, err := pgx.Connect(ctx, dbURL)
		if err != nil {
			return err
		}

		defer db.Close(ctx)

		if err := db.Ping(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to postgres container: %s", err)
	}

	return &DBContainer{
		DSN:      dbURL,
		resource: resource,
	}
}
