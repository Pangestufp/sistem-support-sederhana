package helper

import (
	"TicketManagement/config"
	"TicketManagement/entity"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte(config.ENV.SecretKey)

type JWTClaims struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User, userRoles []entity.UserRole) (string, error) {

	roles := make([]string, len(userRoles))
	for i, r := range userRoles {
		roles[i] = strconv.Itoa(r.RoleID)
	}

	claims := JWTClaims{
		user.ID,
		user.Username,
		roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (*int, []string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, []string{}, errors.New("invalid token signature")
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, []string{}, errors.New("token expired")
		}
		return nil, []string{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, []string{}, errors.New("your token was expired")
	}

	return &claims.ID, claims.Roles, nil
}
