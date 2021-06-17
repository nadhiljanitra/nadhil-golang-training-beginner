package postgres

import (
	"database/sql"
	"strings"

	migrate "github.com/golang-migrate/migrate/v4"
	_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
)

type migration struct {
	Migrate *migrate.Migrate
}

//  Readme https://github.com/golang-migrate/migrate#use-in-your-go-project
func doMigration(db *sql.DB, migrationFolder string) (*migration, error) {
	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationFolder)

	pathToMigrate := strings.Join(dataPath, "")

	driver, err := _postgres.WithInstance(db, &_postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return &migration{Migrate: m}, nil
}

func (m *migration) migrateUp() (bool, error) {
	err := m.Migrate.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

func (m *migration) migrateDown() (bool, error) {
	err := m.Migrate.Down()
	if err != nil {
		return false, err
	}
	return true, err
}
