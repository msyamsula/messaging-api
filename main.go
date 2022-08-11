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

	msgHandler "github.com/msyamsula/messaging-api/domain/message/handler"
	msgService "github.com/msyamsula/messaging-api/domain/message/service"

	"github.com/msyamsula/messaging-api/domain/user/handler"
	"github.com/msyamsula/messaging-api/domain/user/service"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func main() {
	mode := os.Getenv("GIN_MODE")
	fmt.Println(mode, "============")
	if mode != "release" {
		godotenv.Load("env/dev/.env")
	} else {
		godotenv.Load("env/prod/.env")
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
	migrator.Migrate(entity.Message{}, db)

	userService := service.NewService(db)
	messageService := msgService.NewService(db)

	uh := handler.NewHandler(userService)
	mh := msgHandler.NewHandler(messageService)

	allowedOrigins := strings.Split(os.Getenv("ORIGINS"), ",")
	allowedMethods := strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")
	fmt.Println(allowedOrigins)
	fmt.Println(allowedMethods)
	fmt.Println(allowedHeaders)
	CORS := cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           allowedOrigins,
		AllowMethods:           allowedMethods,
		AllowHeaders:           allowedHeaders,
		AllowCredentials:       false,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	})

	// CORS = cors.Default()

	r := gin.Default()
	r.Use(CORS)

	apiPrefix := os.Getenv("API_PREFIX")
	fmt.Println(apiPrefix)

	r.GET(apiPrefix+"/ping", PingHandler)

	r.GET(apiPrefix+"/user/:id", uh.GetSpecificUser)
	r.GET(apiPrefix+"/user", uh.AllUserGet)

	r.POST(apiPrefix+"/register", uh.UserRegister)
	r.POST(apiPrefix+"/login", uh.UserLogin)
	r.POST(apiPrefix+"/logout/:userID", uh.UserLogout)

	r.POST(apiPrefix+"/message", mh.CreateMessage)
	r.POST(apiPrefix+"/message/:senderID", mh.ReadMessages)
	r.GET(apiPrefix+"/message/:senderID/:receiverID", mh.GetMessageByUserID)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	sp := fmt.Sprintf(":%s", port)
	r.Run(sp)
}
