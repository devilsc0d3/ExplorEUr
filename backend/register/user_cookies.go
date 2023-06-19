package register

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var Token = ""

func CreateCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

type JwtClaims struct {
	Nickname string `json:"id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GetTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func CreateJWTToken(nickname, user string) (string, error) {
	jwtSecret := GetEnv("JWT_SECRET")
	var claims = JwtClaims{
		Nickname: nickname,
		Role:     user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tk.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateJWTTokenRememberMe(nickname, user string) (string, error) {
	jwtSecret := GetEnv("JWT_SECRET")
	var claims = JwtClaims{
		Nickname: nickname,
		Role:     user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nil,
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tk.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeJWTToken(token string) (string, string, error) {
	jwtSecret := GetEnv("JWT_SECRET")
	var claims = JwtClaims{}
	tk, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", "", err
	}
	if !tk.Valid {
		return "", "", err
	}
	return claims.Nickname, claims.Role, nil
}
