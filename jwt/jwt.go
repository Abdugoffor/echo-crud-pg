package jwt

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	Secret         string
	Expired        int64
	RefreshExpired int64
}

type IUser[T, E, K any] interface {
	ID() int64
	Pre(http.ResponseWriter, *http.Request, T, ...E) (K, error)
}

type Payload[T IUser[E, K, J], E, K, J any] struct {
	User T `json:"user"`
	jwtv5.RegisteredClaims
}

type JwtService[T IUser[E, K, J], E, K, J any] interface {
	ParseToken(string) (T, error)
	ParseTokenWithExpired(string) (T, error)
	Token(T) (string, error)
}

type jwtService[T IUser[E, K, J], E, K, J any] struct {
	config JwtConfig
}

func (j *jwtService[T, E, K, J]) Token(user T) (string, error) {
	return Encode(user, j.config.Secret, j.config.Expired)
}

func (j *jwtService[T, E, K, J]) ParseToken(token string) (T, error) {
	payload, err := Decode[T, E](token, j.config.Secret, false)
	{
		if err != nil {
			var none T
			return none, err
		}
	}
	return payload.User, nil
}

func (j *jwtService[T, E, K, J]) ParseTokenWithExpired(token string) (T, error) {
	payload, err := Decode[T, E](token, j.config.Secret, true)
	{
		if err != nil {
			var none T
			return none, err
		}
	}
	return payload.User, nil
}

func NewJwtService[T IUser[E, K, J], E, K, J any](config JwtConfig) JwtService[T, E, K, J] {
	return &jwtService[T, E, K, J]{
		config: config,
	}
}

func Encode[T IUser[E, K, J], E, K, J any](user T, secret string, expired int64) (string, error) {

	payload := NewPayload(user, expired)

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(secret))
	{
		if err != nil {
			return "", err
		}
	}

	return tokenString, nil
}

func Decode[T IUser[E, K, J], E, K, J any](tokStr, secret string, withExpired bool) (*Payload[T, E, K, J], error) {

	keyfunc := func(token *jwtv5.Token) (any, error) {
		return []byte(secret), nil
	}

	token, err := jwtv5.ParseWithClaims(tokStr, &Payload[T, E, K, J]{}, keyfunc)
	{
		if err != nil && !errors.Is(err, jwtv5.ErrTokenExpired) {
			return nil, err
		}

		if withExpired && !token.Valid {
			return nil, errors.New("token is not valid")
		}
	}

	payload, ok := token.Claims.(*Payload[T, E, K, J])
	{
		if !ok {
			return nil, errors.New("unknown error")
		}
	}

	return payload, nil
}

func NewPayload[T IUser[E, K, J], E, K, J any](user T, expired int64) *Payload[T, E, K, J] {
	now := time.Now()

	exp := now.Add(time.Minute * time.Duration(expired))

	return &Payload[T, E, K, J]{
		User: user,
		RegisteredClaims: jwtv5.RegisteredClaims{
			ID:        strconv.FormatInt(user.ID(), 10),
			IssuedAt:  jwtv5.NewNumericDate(now),
			ExpiresAt: jwtv5.NewNumericDate(exp),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
		},
	}
}
