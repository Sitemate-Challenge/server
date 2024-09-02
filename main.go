package main

import (
	"fmt"
	"sitemate-challenge-server/internal/db"
	"sitemate-challenge-server/internal/entity"
	issueHandler "sitemate-challenge-server/internal/handler/issue"
	pingHandler "sitemate-challenge-server/internal/handler/ping"
	issueRepo "sitemate-challenge-server/internal/repository/issue"

	"sitemate-challenge-server/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func loadConfig() (*config.Config, error) {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var cfg config.Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal config into struct: %w", err)
	}

	return &cfg, nil
}

func setupDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dbUser := cfg.User
	dbPassword := cfg.Password
	dbName := cfg.Name
	dbHost := cfg.Host
	dbPort := cfg.Port

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	return db.ConnectDB(dsn)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		panic("Failed to setup application config")
	}

	database, err := setupDB(cfg.Database)
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&entity.Issue{})

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Update with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	router := gin.Default()
	router.Use(cors.New(config))

	pingH := pingHandler.New()

	issueRepo := issueRepo.New(database)
	issueHandler := issueHandler.New(issueRepo)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/ping", pingH.Pong)

		apiV1.GET("/issues", issueHandler.GetAllIssues)
		apiV1.GET("/issues/:id", issueHandler.GetIssueByID)
		apiV1.POST("/issues", issueHandler.CreateIssue)
		apiV1.PUT("/issues/:id", issueHandler.UpdateIssue)
		apiV1.DELETE("/issues/:id", issueHandler.DeleteIssue)
	}

	router.Run("localhost:5000")
}
