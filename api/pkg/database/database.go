package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/morelmiles/go-events/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB
}

const (
	maxOpenConns    = 100
	maxIdleConns    = 10
	connMaxLifetime = time.Hour
)

func NewDatabase() (*Database, error) {
	db := &Database{}

	if err := db.initializeDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	if err := db.configureConnectionPool(); err != nil {
		return nil, fmt.Errorf("failed to configure connection pool: %w", err)
	}

	if err := db.runMigrations(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	DB = db.DB
	return db, nil
}

func (d *Database) initializeDB() error {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSL_MODE")

	if sslMode == "" {
		sslMode = "disable"
	}

	dbLink := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=Africa/Nairobi password=%s",
		dbHost, username, dbName, dbPort, sslMode, password)

	var err error
	d.DB, err = gorm.Open(postgres.Open(dbLink), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	log.Println("Database connected successfully🚀")
	return nil
}

func (d *Database) configureConnectionPool() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetConnMaxLifetime(connMaxLifetime)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	return nil
}

func (d *Database) runMigrations() error {
	return d.DB.AutoMigrate(
		&models.User{},
		&models.Event{},
	)
}
