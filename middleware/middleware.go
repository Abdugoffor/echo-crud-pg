package middleware

import (
	"errors"
	"net/http"
	"strings"

	"gitea.avtomig.uz/mydream/shared_service/jwt"
)

type MiddlewareFunc[T any] func(http.ResponseWriter, *http.Request) (T, error)

type AuthMiddleware[T jwt.IUser[E, K, J], E, K, J any] struct {
	jwtService jwt.JwtService[T, E, K, J]
	db         E
}

func NewAuthMiddleware[T jwt.IUser[E, K, J], E, K, J any](
	jwtService jwt.JwtService[T, E, K, J],
	db E,
) *AuthMiddleware[T, E, K, J] {
	return &AuthMiddleware[T, E, K, J]{
		jwtService: jwtService,
		db:         db,
	}
}

func (a *AuthMiddleware[T, E, K, J]) JwtAuthMiddleware(permissions ...K) MiddlewareFunc[J] {

	return func(w http.ResponseWriter, r *http.Request) (J, error) {
		var none J

		token, err := AuthorizationToken(r)
		{
			if err != nil {
				return none, err
			}
		}

		user, err := a.jwtService.ParseTokenWithExpired(token)
		{
			if err != nil {
				return none, err
			}
		}

		data, err := user.Pre(w, r, a.db, permissions...)
		{
			if err != nil {
				return none, err
			}
		}

		return data, nil
	}
}

func AuthorizationToken(r *http.Request) (string, error) {
	const (
		UnAuthorizedErrorMessage = "unauthorized"
		Authorization            = "Authorization"
		Bearer                   = "Bearer "
	)

	token := r.Header.Get(Authorization)
	{
		if token == "" || !strings.HasPrefix(token, Bearer) {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		if strings.Count(token, ".") != 2 {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		token = strings.TrimPrefix(token, Bearer)
	}

	return token, nil
}
