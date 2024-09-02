package main

import (
	"net/http"
	pingHandler "sitemate-challenge-server/internal/handler/ping"
	"sitemate-challenge-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	utils.SendResponse(c, http.StatusOK, "pong", nil)
}

func main() {
	router := gin.Default()

	pingH := pingHandler.New()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/ping", pingH.Pong)
	}

	// issueH := issueHandler.New()
	// dsn := "root:@tcp(127.0.0.1:3306)/interview?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&entity.Issue{})
	// fmt.Println("Database migrated")

	router.Run("localhost:5000")
}
