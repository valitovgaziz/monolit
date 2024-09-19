package storage

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/profclems/go-dotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PSQL_GORM_DB *gorm.DB

func DBconnection() {
	slog.Info("Init DB connection")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		dotenv.GetString("PGHOST"),
		dotenv.GetString("PGUSER"),
		dotenv.GetString("PGPASSWORD"),
		dotenv.GetString("PGDATABASE"),
		dotenv.GetString("PGPORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(2)
	}
	PSQL_GORM_DB = db
	sql, err := db.DB()
	if err != nil {
		slog.Error("failed to get database", "error", err)
		os.Exit(2)
	}
	err = sql.Ping()
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(2)
	}
	slog.Info("DB and DB connections is enabled and ready.")
}
