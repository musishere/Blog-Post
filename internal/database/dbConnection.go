package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/musishere/Blog/config"
	"github.com/musishere/Blog/internal/models"
)

func ConnectDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open("postgres", dsn) // Note: Driver name as first arg in v1
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Post{}, &models.Tag{}, &models.Reaction{})
	if db.Error != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", db.Error)
	}

	return db, nil // Return the actual connected DB, not a new empty one
}
