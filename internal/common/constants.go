package common

import "time"

// Error Code
const (
	ErrorInvalidParameters = "INVALID_PARAMETERS"
	ErrorInternalError = "INTERNAL_ERROR"
	ErrorSslBackendError = "SSL_BACKEND_ERROR"
	ErrorGtsError = "SSL_GTS_ERROR"
)

// Empty Value
var EmptyMap = map[string]interface{}{}

var timeLayout = "2006-01-02T15:04:05.000Z"
var timeStr = "2099-12-31T23:59:59.991Z"
var DefaultExpiration, _ = time.Parse(timeLayout, timeStr)


