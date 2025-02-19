package request

import "git.sriss.uz/shared/shared_service/response"

func (r *request) OK(data ...any) error {
	return response.Response(r.ctx).OK(data...)
}

func (r *request) Created(data ...any) error {
	return response.Response(r.ctx).Created(data...)
}

func (r *request) Accepted(data ...any) error {
	return response.Response(r.ctx).Accepted(data...)
}

func (r *request) NonAuthoritativeInfo(data ...any) error {
	return response.Response(r.ctx).NonAuthoritativeInfo(data...)
}

func (r *request) NoContent() error {
	return response.Response(r.ctx).NoContent()
}

func (r *request) ResetContent(data ...any) error {
	return response.Response(r.ctx).ResetContent(data...)
}

func (r *request) PartialContent(data ...any) error {
	return response.Response(r.ctx).PartialContent(data...)
}

func (r *request) MultiStatus(data ...any) error {
	return response.Response(r.ctx).MultiStatus(data...)
}

func (r *request) AlreadyReported(data ...any) error {
	return response.Response(r.ctx).AlreadyReported(data...)
}

func (r *request) IMUsed(data ...any) error {
	return response.Response(r.ctx).IMUsed(data...)
}
