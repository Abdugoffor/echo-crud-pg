package response

import "net/http"

type HttpSuccess interface {
	// 200 - OK
	OK(data ...any) error
	// 201 - Created
	Created(data ...any) error
	// 202 - Accepted
	Accepted(data ...any) error
	// 203 - Non-Authoritative Information
	NonAuthoritativeInfo(data ...any) error
	// 204 - No Content
	NoContent() error
	// 205 - Reset Content
	ResetContent(data ...any) error
	// 206 - Partial Content
	PartialContent(data ...any) error
	// 207 - Multi-Status
	MultiStatus(data ...any) error
	// 208 - Already Reported
	AlreadyReported(data ...any) error
	// 226 - IM Used
	IMUsed(data ...any) error
}

var _ HttpSuccess = (*response)(nil)

func (r *response) OK(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusOK)
	}

	return r.ctx.JSON(http.StatusOK, data)
}

func (r *response) Created(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusCreated)
	}

	return r.ctx.JSON(http.StatusCreated, data)
}

func (r *response) Accepted(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusAccepted)
	}

	return r.ctx.JSON(http.StatusAccepted, data)
}

func (r *response) NonAuthoritativeInfo(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	return r.ctx.JSON(http.StatusNonAuthoritativeInfo, data)
}

func (r *response) NoContent() error {
	return r.ctx.NoContent(http.StatusNoContent)
}

func (r *response) ResetContent(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusResetContent)
	}

	return r.ctx.JSON(http.StatusResetContent, data)
}

func (r *response) PartialContent(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusPartialContent)
	}

	return r.ctx.JSON(http.StatusPartialContent, data)
}

func (r *response) MultiStatus(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusMultiStatus)
	}

	return r.ctx.JSON(http.StatusMultiStatus, data)
}

func (r *response) AlreadyReported(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusAlreadyReported)
	}

	return r.ctx.JSON(http.StatusAlreadyReported, data)
}

func (r *response) IMUsed(data ...any) error {

	if len(data) > 0 {
		return r.ctx.NoContent(http.StatusIMUsed)
	}

	return r.ctx.JSON(http.StatusIMUsed, data)
}
