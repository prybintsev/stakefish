package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/prybintsev/stakefish/internal/config"
)

func ConnectToPostgres(cfg config.AppConfig) *sql.DB {
	var connStr string
	if cfg.DBUser != "" && cfg.DBPassword != "" {
		connStr = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable", cfg.DBHost,
			cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBPort)
	} else {
		connStr = fmt.Sprintf("host=%s dbname=%s port=%d sslmode=disable", cfg.DBHost,
			cfg.DBName, cfg.DBPort)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
