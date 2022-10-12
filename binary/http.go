package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	msgDBI "github.com/msyamsula/messaging-api/message/database"
	msgDB "github.com/msyamsula/messaging-api/message/database/object"
	msgHandler "github.com/msyamsula/messaging-api/message/handler/object"
	msgSvcI "github.com/msyamsula/messaging-api/message/service"
	msgSvc "github.com/msyamsula/messaging-api/message/service/object"

	userDBI "github.com/msyamsula/messaging-api/user/database"
	userDB "github.com/msyamsula/messaging-api/user/database/object"
	userHandler "github.com/msyamsula/messaging-api/user/handler/object"
	userSvcI "github.com/msyamsula/messaging-api/user/service"
	userSvc "github.com/msyamsula/messaging-api/user/service/object"

	mdwareHandlerO "github.com/msyamsula/messaging-api/middleware/handler/object"
	tokenO "github.com/msyamsula/messaging-api/middleware/token/object"
)

func main() {

	var userService userSvcI.Service
	{
		cfg := userDBI.DbConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USERNAME"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME"),
		}
		db, err := userDB.New(cfg)
		if err != nil {
			panic(err)
		}

		svc, err := userSvc.New(db)
		if err != nil {
			panic(err)
		}

		userService = svc

	}

	var messageService msgSvcI.Service
	{
		mongoURI := os.Getenv("MONGO_URI")
		var err error
		var db msgDBI.Database

		db, err = msgDB.New(mongoURI)
		if err != nil {
			panic(err)
		}

		messageService = msgSvc.New(db)
	}

	secret := os.Getenv("JSON_SECRET")
	expiryTime := 30 * time.Minute
	tokenI := tokenO.New([]byte(secret), expiryTime)

	r := gin.Default()
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
	r.Use(CORS)

	// handler
	mdwareHandlerI := mdwareHandlerO.New(tokenI)
	userHandler := userHandler.New(userService, tokenI)
	msgHandler := msgHandler.New(messageService)

	apiPrefix := os.Getenv("API_PREFIX")
	r.GET(apiPrefix+"/ping", userHandler.Pong)
	r.GET(apiPrefix+"/login", userHandler.Login)
	r.POST(apiPrefix+"/register", userHandler.Register)
	r.GET(apiPrefix+"/users", mdwareHandlerI.ValidateToken, userHandler.GetAllUser)
	r.POST(apiPrefix+"/logout", userHandler.Logout)

	r.POST(apiPrefix+"/message", mdwareHandlerI.ValidateToken, msgHandler.InsertMessage)
	r.GET(apiPrefix+"/message", mdwareHandlerI.ValidateToken, msgHandler.GetConversation)
	r.PUT(apiPrefix+"/message", mdwareHandlerI.ValidateToken, msgHandler.ReadMessage)

	r.Run(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	database "github.com/msyamsula/messaging-api/db"
// 	"github.com/msyamsula/messaging-api/db/entity"
// 	"github.com/msyamsula/messaging-api/db/migrator"

// 	msgHandler "github.com/msyamsula/messaging-api/domain/message/handler"
// 	msgService "github.com/msyamsula/messaging-api/domain/message/service"

// 	"github.com/msyamsula/messaging-api/domain/user/handler"
// 	"github.com/msyamsula/messaging-api/domain/user/service"
// )

// func PingHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, "pong")
// }

// func main() {
// 	godotenv.Load(".env")
// 	mode := os.Getenv("GIN_MODE")
// 	fmt.Println(mode, "============")
// 	if mode == "release" {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	config := database.PgConfig{
// 		Host:     os.Getenv("DB_HOST"),
// 		Port:     os.Getenv("DB_PORT"),
// 		Dbname:   os.Getenv("DB_NAME"),
// 		User:     os.Getenv("DB_USERNAME"),
// 		Password: os.Getenv("DB_PASSWORD"),
// 	}

// 	fmt.Println(config)
// 	database.Connect(config)
// 	db := database.GetDB()

// 	migrator.Migrate(entity.User{}, db)
// 	migrator.Migrate(entity.Message{}, db)

// 	userService := service.NewService(db)
// 	messageService := msgService.NewService(db)

// 	uh := handler.NewHandler(userService)
// 	mh := msgHandler.NewHandler(messageService)

// 	// CORS = cors.Default()

// 	r := gin.Default()

// 	apiPrefix := os.Getenv("API_PREFIX")
// 	fmt.Println(apiPrefix)

// 	r.GET(apiPrefix+"/ping", PingHandler)

// 	r.GET(apiPrefix+"/user/:id", uh.GetSpecificUser)
// 	r.GET(apiPrefix+"/user", uh.AllUserGet)

// 	r.POST(apiPrefix+"/register", uh.UserRegister)
// 	r.POST(apiPrefix+"/login", uh.UserLogin)
// 	r.POST(apiPrefix+"/logout/:userID", uh.UserLogout)

// 	r.POST(apiPrefix+"/message", mh.CreateMessage)
// 	r.POST(apiPrefix+"/message/:senderID", mh.ReadMessages)
// 	r.GET(apiPrefix+"/message/:senderID/:receiverID", mh.GetMessageByUserID)

// 	port := os.Getenv("APP_PORT")
// 	fmt.Println(port)
// 	sp := fmt.Sprintf(":%s", port)
// 	r.Run(sp)
// }
