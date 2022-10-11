package main

import (
	"fmt"
	"os"
	"time"

	tokenObj "github.com/msyamsula/messaging-api/middleware/token/object"
)

func main() {

	secretStr := os.Getenv("JSON_SECRET")
	fmt.Println(secretStr)
	expDuration := 10 * time.Minute
	token := tokenObj.New([]byte(secretStr), expDuration)

	userID := 2
	t, err := token.Create(int64(userID))
	fmt.Println(t)
	// var jt *jwt.Claims
	t, err = token.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjU1MDEyNzgsInVzZXJJRCI6Mn0.r9xrJlZ7YnJxBch9FOTc2bmfWW8c6TGbW-fOztyAcPM")
	// fmt.Println(t)
	fmt.Println(err, t)

	// var err error
	// var db database.Database

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// m := database.MessageToInsert{
	// 	SenderID:   3,
	// 	Text:       "mantap",
	// 	ReceiverID: 5,
	// 	IsRead:     false,
	// 	CreatedAt:  time.Now().Unix(),
	// }

	// var messages []database.Message
	// err = messageService.InsertMessage(ctx, m)
	// err = messageService.ReadMessage(ctx, 3, 5)
	// messages, err = messageService.GetConversation(ctx, 5, 3)

	// fmt.Println(messages)
	// fmt.Println(err)

}
