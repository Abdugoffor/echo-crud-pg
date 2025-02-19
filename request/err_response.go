package request

import "git.sriss.uz/shared/shared_service/response"

func (r *request) BadRequest(err ...any) error {
	return response.HTTPError(err...).BadRequest()
}

func (r *request) Unauthorized(err ...any) error {
	return response.HTTPError(err...).Unauthorized()
}

func (r *request) PaymentRequired(err ...any) error {
	return response.HTTPError(err...).PaymentRequired()
}

func (r *request) Forbidden(err ...any) error {
	return response.HTTPError(err...).Forbidden()
}

func (r *request) NotFound(err ...any) error {
	return response.HTTPError(err...).NotFound()
}

func (r *request) MethodNotAllowed(err ...any) error {
	return response.HTTPError(err...).MethodNotAllowed()
}

func (r *request) NotAcceptable(err ...any) error {
	return response.HTTPError(err...).NotAcceptable()
}

func (r *request) ProxyAuthRequired(err ...any) error {
	return response.HTTPError(err...).ProxyAuthRequired()
}

func (r *request) RequestTimeout(err ...any) error {
	return response.HTTPError(err...).RequestTimeout()
}

func (r *request) Conflict(err ...any) error {
	return response.HTTPError(err...).Conflict()
}

func (r *request) Gone(err ...any) error {
	return response.HTTPError(err...).Gone()
}

func (r *request) LengthRequired(err ...any) error {
	return response.HTTPError(err...).LengthRequired()
}

func (r *request) PreconditionFailed(err ...any) error {
	return response.HTTPError(err...).PreconditionFailed()
}

func (r *request) RequestEntityTooLarge(err ...any) error {
	return response.HTTPError(err...).RequestEntityTooLarge()
}

func (r *request) RequestURITooLong(err ...any) error {
	return response.HTTPError(err...).RequestURITooLong()
}

func (r *request) UnsupportedMediaType(err ...any) error {
	return response.HTTPError(err...).UnsupportedMediaType()
}

func (r *request) RequestedRangeNotSatisfiable(err ...any) error {
	return response.HTTPError(err...).RequestedRangeNotSatisfiable()
}

func (r *request) SessionExpired(err ...any) error {
	return response.HTTPError(err...).SessionExpired()
}

func (r *request) ExpectationFailed(err ...any) error {
	return response.HTTPError(err...).ExpectationFailed()
}

func (r *request) MisdirectedRequest(err ...any) error {
	return response.HTTPError(err...).MisdirectedRequest()
}

func (r *request) UnprocessableEntity(err ...any) error {
	return response.HTTPError(err...).UnprocessableEntity()
}

func (r *request) Locked(err ...any) error {
	return response.HTTPError(err...).Locked()
}

func (r *request) FailedDependency(err ...any) error {
	return response.HTTPError(err...).FailedDependency()
}

func (r *request) TooEarly(err ...any) error {
	return response.HTTPError(err...).TooEarly()
}

func (r *request) UpgradeRequired(err ...any) error {
	return response.HTTPError(err...).UpgradeRequired()
}

func (r *request) PreconditionRequired(err ...any) error {
	return response.HTTPError(err...).PreconditionRequired()
}

func (r *request) TooManyRequests(err ...any) error {
	return response.HTTPError(err...).TooManyRequests()
}

func (r *request) RequestHeaderFieldsTooLarge(err ...any) error {
	return response.HTTPError(err...).RequestHeaderFieldsTooLarge()
}

func (r *request) UnavailableForLegalReasons(err ...any) error {
	return response.HTTPError(err...).UnavailableForLegalReasons()
}

func (r *request) InternalServerError(err ...any) error {
	return response.HTTPError(err...).InternalServerError()
}

func (r *request) NotImplemented(err ...any) error {
	return response.HTTPError(err...).NotImplemented()
}

func (r *request) BadGateway(err ...any) error {
	return response.HTTPError(err...).BadGateway()
}

func (r *request) ServiceUnavailable(err ...any) error {
	return response.HTTPError(err...).ServiceUnavailable()
}

func (r *request) GatewayTimeout(err ...any) error {
	return response.HTTPError(err...).GatewayTimeout()
}

func (r *request) HTTPVersionNotSupported(err ...any) error {
	return response.HTTPError(err...).HTTPVersionNotSupported()
}

func (r *request) VariantAlsoNegotiates(err ...any) error {
	return response.HTTPError(err...).VariantAlsoNegotiates()
}

func (r *request) InsufficientStorage(err ...any) error {
	return response.HTTPError(err...).InsufficientStorage()
}

func (r *request) LoopDetected(err ...any) error {
	return response.HTTPError(err...).LoopDetected()
}

func (r *request) NotExtended(err ...any) error {
	return response.HTTPError(err...).NotExtended()
}

func (r *request) NetworkAuthenticationRequired(err ...any) error {
	return response.HTTPError(err...).NetworkAuthenticationRequired()
}
