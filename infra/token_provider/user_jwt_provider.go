package tokenprovider

import (
	"app/configs"
	"app/internal/user/entity"
	"app/internal/user/param/response"
	"app/internal/user/port/driven"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	_ driven.TokenProvider[*entity.User] = new(UserJwtProvider)
)

type UserJwtProvider struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	ExpiresSecond int
}

func NewUserJwtProvider(conf *configs.ApplicationConfig) *UserJwtProvider {
	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(conf.JWT.PrivateKey))
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(conf.JWT.PublicKey))
	if err != nil {
		panic(err)
	}
	return &UserJwtProvider{
		PrivateKey:    secretKey,
		PublicKey:     publicKey,
		ExpiresSecond: conf.JWT.ExpiresSecond,
	}
}

func (utp *UserJwtProvider) Generate(user *entity.User) (*response.Token, error) {
	jwtID, _ := uuid.NewRandom()
	claims := jwt.RegisteredClaims{
		Issuer:    "dating-be",
		Subject:   fmt.Sprintf("%d", user.ID),
		Audience:  []string{"dating-be"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(utp.ExpiresSecond))),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        jwtID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(utp.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &response.Token{
		Token:     tokenString,
		ExpiresIn: int(time.Hour.Seconds()),
		Type:      "Bearer",
	}, nil
}
