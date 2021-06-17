package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Suite struct {
	suite.Suite

	Username                string
	Password                string
	DBName                  string
	MigrationLocationFolder string

	container testcontainers.Container
	db        *sql.DB
	migration *migration
}

func NewSuite(username string, password string, dbName string, migrationDir string) *Suite {
	return &Suite{
		Username:                username,
		Password:                password,
		DBName:                  dbName,
		MigrationLocationFolder: migrationDir,
	}
}

func NewDefaultSuite(migrationDir string) *Suite {
	return NewSuite("user", "password", "postgres_test", migrationDir)
}

func (s *Suite) SetupSuite() {
	ctx := context.TODO()

	// prepare env for test container
	env := map[string]string{
		"POSTGRES_DB":       s.DBName,
		"POSTGRES_USER":     s.Username,
		"POSTGRES_PASSWORD": s.Password,
	}

	// request to create new postgres container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:11.4-alpine",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		Env:          env,
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	// generate posrgres container
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	s.Require().NoError(err)
	s.container = postgresC

	// get Port from the container
	port, err := postgresC.MappedPort(ctx, "5432/tcp")
	s.Require().NoError(err)

	// construct sql data source name and open connection
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=localhost port=%d sslmode=disable", s.Username, s.Password, s.DBName, port.Int())
	s.db, err = sql.Open("postgres", dsn)
	s.Require().NoError(err)

	err = s.db.Ping()
	s.Require().NoError(err)

	s.migration, err = doMigration(s.db, s.MigrationLocationFolder)
	s.Require().NoError(err)
}

// closing db connection after done testing
func (s *Suite) TearDownSuite() {
	defer func() {
		_ = s.container.Terminate(context.Background())
	}()

	err := s.db.Close()
	s.Require().NoError(err)
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	ok, err := s.migration.migrateUp()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s *Suite) AfterTest(suiteName, testName string) {
	ok, err := s.migration.migrateDown()
	s.Require().NoError(err)
	s.Require().True(ok)
}
