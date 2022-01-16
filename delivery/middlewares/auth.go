package middleware

import (
	"errors"
	"strings"
	"time"
	"todolist-app/configs"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Auth interface {
	GenerateToken(userID int) (string, error)
	ExtractTokenUserID(e echo.Context) int
}

type jwtService struct {
}

var SecretKey = configs.SecretKey

func NewAuth() *jwtService {
	return &jwtService{}
}

func (a *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Minute * 5).Unix()
	// claim["exp"] = time.Now().AddDate(0, 0, 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (a *jwtService) ExtractTokenUserID(c echo.Context) int {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

	if !strings.Contains(authHeader, "Bearer") {
		return 0
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SecretKey), nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0
	}

	userID := int(claim["user_id"].(float64))

	return userID

}
