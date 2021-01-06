package common

// Error Code
const (
	ErrorInvalidParameters = "INVALID_PARAMETERS"
	ErrorInternalError = "INTERNAL_ERROR"
	ErrorSslBackendError = "SSL_BACKEND_ERROR"
	ErrorGtsError = "SSL_GTS_ERROR"
)

// Empty Value
var EmptyMap = map[string]interface{}{}

// Domain related constants
const (
	DomainCertStatusPending = "Pending"
	DomainCertStatusDone = "Done"
	DomainCertStatusGts = "GTS"
)
