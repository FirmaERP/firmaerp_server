package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func nowFunc() time.Time {
	return time.Now().UTC()
}

func InitGorm() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("the environment variable DB_HOST is empty")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, errors.New("the environment variable DB_PORT is empty")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("the environment variable DB_NAME is empty")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("the environment variable DB_USER is empty")
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		return nil, errors.New("the environment variable DB_PASS is empty")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		host,
		port,
		name,
		user,
		pass,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NowFunc: nowFunc})
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to the database")
	}

	return db, nil
}
