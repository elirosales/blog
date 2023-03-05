package testutils

import (
	"testing"

	"github.com/elizabethrosales/blog/config"
	"github.com/elizabethrosales/blog/database"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type TestSuite struct {
	T      *testing.T
	DB     *gorm.DB
	Config config.Config
	Seed   string
}

func NewSuite(t *testing.T, seed string, configFile string) *TestSuite {
	log := logrus.New()

	err := godotenv.Load(configFile)
	if err != nil {
		log.Errorf("Failed to load config file")
	}

	c := config.New()

	testDB := NewTestDB()
	c.Database.DSN = testDB.Postgres.DSN

	db, _ := database.Initialize(*c)

	return &TestSuite{
		T:      t,
		DB:     db,
		Config: *c,
		Seed:   seed,
	}
}

func (ts *TestSuite) Run(t *testing.T, name string, f func(t *testing.T)) {
	t.Run(name, func(t *testing.T) {
		ts.SetupTest()
		defer ts.TearDownTest()
		f(t)
	})
}

func (ts *TestSuite) SetupTest() {
	if ts.Seed == "" {
		return
	}

	require.NoError(ts.T, ts.createTables())
	if ts.Seed != "" {
		ts.seedPostgres()
	}
}

func (ts *TestSuite) TearDownTest() {
	if ts.Seed != "" {
		ts.clearTable()
	}
}
