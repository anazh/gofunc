package http_resp

// define http status for http response
// we just define common http status

const (
	HTTP_STATUS_OK                            = 200 // when server responded ok for http request
	HTTP_STATUS_NON_AUTHORITATIVE_INFORMATION = 203 // when server not get auth information,but is necessary
	HTTP_STATUS_MOVED_PERMANENTLY             = 301 // redirect url
	HTTP_STATUS_FOUND                         = 302 // redirect url
	HTTP_STATUS_BAD_REQUEST                   = 400 // http request is invalid
	HTTP_STATUS_UNAUTHORIZED                  = 401 // http request is not authorized
	HTTP_STATUS_FORBIDDEN                     = 403 // http request is forbidden
	HTTP_STATUS_NOT_FOUND                     = 404 // http request is not found
	HTTP_STATUS_METHOD_NOT_ALLOWED            = 405 // http request method is not allowed
	HTTP_STATUS_REQUEST_TIMEOUT               = 408 // http request time out
	HTTP_STATUS_CONFLICT                      = 409 // http request conflict
	HTTP_STATUS_LENGTH_REQUIRED               = 411 // http request length required
	HTTP_STATUS_REQUEST_ENTITY_TOO_LARGE      = 413 // http request entity too large
	HTTP_STATUS_UNSUPPORTED_MEDIA_TYPE        = 415 // http request unsupported media type
	HTTP_STATUS_INTERNAL_SERVER_ERROR         = 500 // http server internal server error
	HTTP_STATUS_BAD_GATEWAY                   = 502 // bad gateway
	HTTP_STATUS_SERVICE_UNAVAILABLE           = 503 // service unavailable
	HTTP_STATUS_GATEWAY_TIMEOUT               = 504 // server time out
)
