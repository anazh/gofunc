package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errTokenFormat = errors.New("token format is error")
)

type TokenInfo struct {
	Sign    []byte //sign to encode token,and decode
	UserId  uint
	Data    string
	ExpTime int64 //token expiration time,use second as unit of time
}

func GenToken(t TokenInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   t.UserId,
		"data":      t.Data,
		"sign_time": time.Now().Unix(),
		"exp_time":  t.ExpTime,
	})
	return token.SignedString(t.Sign) // Sign and get the complete encoded token as a string using the secret
}

// decode string to token info
func DecodeToken(tokenString string, sign []byte) (TokenInfo, error) {
	if tokenString == "" {
		return TokenInfo{}, errors.New("token is empty")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Don't forget to validate the alg is what you expect:
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return sign, nil // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	})
	if err != nil {
		return TokenInfo{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expTime, ok := claims["exp_time"].(int64)
		if ok && time.Now().Unix() < expTime {
			out := TokenInfo{}
			out.UserId = claims["user_id"].(uint)
			out.Data = claims["data"].(string)
			out.ExpTime = expTime
			return out, nil
		}
		return TokenInfo{}, errors.New("invalid token,token timeout")
	}
	return TokenInfo{}, err
}

// get token form head string
// format: Bearer token
func TokenStrFromHeader(tokenStr string) (string, error) {
	if tokenStr[0:6] == "Bearer" {
		tokenStr = tokenStr[7:]
		return tokenStr, nil
	}
	return "", errTokenFormat
}
