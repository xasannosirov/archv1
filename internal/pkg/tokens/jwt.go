package tokens

import (
	"archv1/internal/pkg/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

// JWTHandler ...
type JWTHandler struct {
	Sub        int
	Iss        string
	Exp        string
	Iat        string
	Aud        []string
	Role       string
	SigningKey string
	Log        *log.Logger
	Token      string
	Timeout    int
}

// CustomClaims ...
type CustomClaims struct {
	*jwt.Token
	Sub  int     `json:"sub"`
	Exp  float64 `json:"exp"`
	Iat  float64 `json:"iat"`
	Role string  `json:"role"`
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
		rtClaims     jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHandler.Sub
	claims["exp"] = cast.ToInt(time.Now().Unix()) + jwtHandler.Timeout
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role

	cfg := config.NewConfig()

	access, err = accessToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		log.Println("errors generating access token: ", err)
		return
	}

	rtClaims = refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = jwtHandler.Sub
	rtClaims["exp"] = cast.ToInt(time.Now().Unix()) + cast.ToInt(cfg.RefreshTTL)
	rtClaims["iat"] = time.Now().Unix()
	rtClaims["role"] = jwtHandler.Role

	refresh, err = refreshToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		log.Println("errors generating refresh token: ", err)
		return
	}

	return
}

// ExtractClaims ...
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	cfg := config.NewConfig()

	token, err = jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		log.Println("invalid jwt token")
		return nil, err
	}

	return claims, nil
}

// ExtractClaim extracts claims from given token
func ExtractClaim(tokenStr string, signingKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = errors.New("invalid JWT Token")
		return nil, err
	}

	return claims, nil
}
