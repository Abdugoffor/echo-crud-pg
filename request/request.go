package request

import (
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	UnAuthorizedErrorMessage = "unauthorized"
	Authorization            = "Authorization"
	Bearer                   = "Bearer "
)

func Request(ctx echo.Context) *request {
	return &request{
		ctx:    ctx,
		binder: echo.DefaultBinder{},
	}
}

type request struct {
	ctx    echo.Context
	binder echo.DefaultBinder
}

func (r *request) BindParam(in any) error {
	return r.binder.BindPathParams(r.ctx, in)
}

func (r *request) BindQuery(in any) error {
	return r.binder.BindQueryParams(r.ctx, in)
}

func (r *request) BindHeader(in any) error {
	return r.binder.BindHeaders(r.ctx, in)
}

func (r *request) BindBody(in any) error {
	return r.binder.BindBody(r.ctx, in)
}

func (r *request) Bind(in any) error {
	return r.binder.Bind(in, r.ctx)
}

func (r *request) Param(name string) string {
	return r.ctx.Param(name)
}

func (r *request) Query(name string) string {
	return r.ctx.QueryParam(name)
}

func (r *request) Header(name string) string {
	return r.ctx.Request().Header.Get(name)
}

func (r *request) ParamToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Param(name), 10, 64)
}

func (r *request) QueryToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Query(name), 10, 64)
}

func (r *request) HeaderToInt(name string) (int64, error) {
	return strconv.ParseInt(r.Header(name), 10, 64)
}

func (r *request) ParamToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Param(name), 64)
}

func (r *request) QueryToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Query(name), 64)
}

func (r *request) HeaderToFloat(name string) (float64, error) {
	return strconv.ParseFloat(r.Header(name), 64)
}

func (r *request) ParamToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Param(name))
}

func (r *request) QueryToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Query(name))
}

func (r *request) HeaderToBool(name string) (bool, error) {
	return strconv.ParseBool(r.Header(name))
}

func (r *request) NewPaginate() *Paginate {
	return NewPaginateEchoWithContext(r.ctx)
}

func (r *request) AuthorizationToken() string {
	return r.Header(Authorization)
}

func (r *request) AuthorizationTokenWithBearer() (string, error) {

	token := r.AuthorizationToken()
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
