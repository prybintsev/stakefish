package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/config"
)

func MigrateUp(db *sql.DB, cfg config.AppConfig, logEntry *logrus.Entry) error {
	driver, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return fmt.Errorf("cannot create Postgres driver: %v", err)
	}
	path, _ := os.Getwd()
	logrus.Info(fmt.Sprintf("current dir is %s", path))
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", cfg.MigrationsPath), cfg.DBName, driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logEntry.Info("Migrations: no change")
			return nil
		}
		return err
	}

	logEntry.Info("Migrations have completed successfully")

	return nil
}
