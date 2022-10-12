package object

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/msyamsula/messaging-api/middleware/handler"
	tokenI "github.com/msyamsula/messaging-api/middleware/token"
)

type JWT struct {
	Secret      []byte
	ExpDuration time.Duration
}

func New(s []byte, d time.Duration) tokenI.Token {
	return &JWT{
		Secret:      s,
		ExpDuration: d,
	}
}

func (j *JWT) Create(userID int64) (string, error) {
	var err error
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(j.ExpDuration).Unix()
	claims["userID"] = userID

	var tokenStr string
	tokenStr, err = token.SignedString(j.Secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (j *JWT) Validate(token string) (string, error) {
	if token == "" {
		return token, handler.ErrNoToken
	}

	var err error
	// var t *jwt.Token
	_, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", tokenI.ErrSignAlgoNotMatched
		}
		return j.Secret, nil
	})

	return token, err
}
