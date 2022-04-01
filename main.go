package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/msyamsula/messaging-api/db"
	"github.com/msyamsula/messaging-api/db/entity"
	"github.com/msyamsula/messaging-api/db/migrator"

	"github.com/msyamsula/messaging-api/user/handler"
	"github.com/msyamsula/messaging-api/user/service"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func main() {
	mode := os.Getenv("GIN_MODE")
	fmt.Println(mode, "============")
	if mode != "release" {
		godotenv.Load("dev/.env")
	} else {
		fmt.Println("goes here")
		godotenv.Load("prod/.env")
	}

	config := database.PgConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	fmt.Println(config)
	database.Connect(config)
	db := database.GetDB()

	migrator.Migrate(entity.User{}, db)
	userService := service.NewService(db)

	uh := handler.NewHandler(userService)

	allowedOrigins := strings.Split(os.Getenv("ORIGINS"), ",")

	CORS := cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           allowedOrigins,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:           []string{"content-type"},
		AllowCredentials:       false,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	})

	r := gin.Default()
	r.Use(CORS)

	r.GET("/ping", PingHandler)

	r.GET("/user/:id", uh.GetSpecificUser)
	r.GET("/user", uh.AllUserGet)

	r.POST("/register", uh.UserRegister)
	r.POST("/login", uh.UserLogin)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	sp := fmt.Sprintf(":%s", port)
	r.Run(sp)
}
