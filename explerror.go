package explerror

import (
	"encoding/json"
	"log"
	"net/http"
)

var logger *log.Logger = nil
var send func(w http.ResponseWriter, status int, data *Error) error = func(w http.ResponseWriter, status int, data *Error) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// Options represents a configuration container to setup the Explerror package.
type Options struct {
	// Debug enables logging to the specified Logger.
	// Uses log.Default() if no logger is specified.
	Debug bool

	// Log is a logger which can be used to output errors.
	Log *log.Logger

	// SendFunction is a function which gets used to send the error to the response writer.
	SendFunction func(w http.ResponseWriter, status int, data *Error) error
}

// New creates a new explerror instance with the provided options.
func New(options Options) {
	if options.Log != nil {
		logger = options.Log
	}

	if options.Debug && options.Log != nil {
		logger = log.Default()
	}

	// send defaults if no sendFunction is provided
	if options.SendFunction != nil {
		send = options.SendFunction
	}
}

// Default create a new explerror instance with the default options.
func Default() {
	New(Options{
		Debug:        false,
		Log:          nil,
		SendFunction: nil,
	})
}

// sendError sends an error message using the provided send-function and logs the error
func sendError(w http.ResponseWriter, statusCode int, err error) {
	logger.Printf("%d: %s", statusCode, err.Error())

	theError := &Error{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	send(w, statusCode, theError)
}

// VariantAlsoNegotiates sends an error with status 506 Variant Also Negotiates
func VariantAlsoNegotiates(w http.ResponseWriter, err error) {
	sendError(w, http.StatusVariantAlsoNegotiates, err)
}

// UpgradeRequired sends an error with status 426 Ugrade Required
func UpgradeRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUpgradeRequired, err)
}

// UnsupportedMediaType sends an error with status 415 Unsupported Media Type
func UnsupportedMediaType(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnsupportedMediaType, err)
}

// UnavailableForLegalReasons sends an error with status 451 Unavailable For Legal Reasons
func UnavailableForLegalReasons(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnavailableForLegalReasons, err)
}

// ServiceUnavailable sends an error with status 503 Service Unavailable
func ServiceUnavailable(w http.ResponseWriter, err error) {
	sendError(w, http.StatusServiceUnavailable, err)
}

// RequestedRangeNotSatisfiable sends an error with status 416 Requested Range Not Satisfiable
func RequestedRangeNotSatisfiable(w http.ResponseWriter, err error) {
	sendError(w, http.StatusRequestedRangeNotSatisfiable, err)
}

// RequestURITooLong sends an error with status 414 Request URI Too Long
func RequestURITooLong(w http.ResponseWriter, err error) {
	sendError(w, http.StatusRequestURITooLong, err)
}

// RequestHeaderFieldsTooLarge sends an error with status 431 RequestHeaderFieldsTooLarge
func RequestHeaderFieldsTooLarge(w http.ResponseWriter, err error) {
	sendError(w, http.StatusRequestHeaderFieldsTooLarge, err)
}

// RequestEntityTooLarge sends an error with status 413 Request Entity Too Large
func RequestEntityTooLarge(w http.ResponseWriter, err error) {
	sendError(w, http.StatusRequestEntityTooLarge, err)
}

// TooEarly sends an error with status 425 Too Early
func TooEarly(w http.ResponseWriter, err error) {
	sendError(w, http.StatusTooEarly, err)
}

// ProxyAuthRequired sends an error with status 407 Proxy Auth Required
func ProxyAuthRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusProxyAuthRequired, err)
}

// PreconditionRequired sends an error with status 428 Precondition Required
func PreconditionRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusPreconditionRequired, err)
}

// PreconditionFailed sends an error with status 412 Precondition Failed
func PreconditionFailed(w http.ResponseWriter, err error) {
	sendError(w, http.StatusPreconditionFailed, err)
}

// PaymentRequired sends an error with status 402 Payment Required
func PaymentRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusPaymentRequired, err)
}

// NotImplemented sends an error with status 501 Not Implemented
func NotImplemented(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotImplemented, err)
}

// NotExtended sends an error with status 510 Not Extended
func NotExtended(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotExtended, err)
}

// NotAcceptable sends an error with status 406 Not Acceptable
func NotAcceptable(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotAcceptable, err)
}

// NetworkAuthenticationRequired sends an error with status 511 Network Authentication Required
func NetworkAuthenticationRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNetworkAuthenticationRequired, err)
}

// MethodNotAllowed sends an error with status 405 Method Not Allowed
func MethodNotAllowed(w http.ResponseWriter, err error) {
	sendError(w, http.StatusMethodNotAllowed, err)
}

// LoopDetected sends an error with status 508 Loop Detected
func LoopDetected(w http.ResponseWriter, err error) {
	sendError(w, http.StatusLoopDetected, err)
}

// Locked sends an error with status 423 Locked
func Locked(w http.ResponseWriter, err error) {
	sendError(w, http.StatusLocked, err)
}

// LengthRequired sends an error with status 411 Length Required
func LengthRequired(w http.ResponseWriter, err error) {
	sendError(w, http.StatusLengthRequired, err)
}

// InsufficientStorage sends an error with status 507 Insufficient Storage
func InsufficientStorage(w http.ResponseWriter, err error) {
	sendError(w, http.StatusInsufficientStorage, err)
}

// HTTPVersionNotSupported sends an error with status 505 HTTP Version Not Supported
func HTTPVersionNotSupported(w http.ResponseWriter, err error) {
	sendError(w, http.StatusHTTPVersionNotSupported, err)
}

// Gone sends an error with status 410 Gone
func Gone(w http.ResponseWriter, err error) {
	sendError(w, http.StatusGone, err)
}

// FailedDependency sends an error with status 424 Failed Dependency
func FailedDependency(w http.ResponseWriter, err error) {
	sendError(w, http.StatusFailedDependency, err)
}

// GatewayTimeout sends an error with status 504 Gateway Timeout
func GatewayTimeout(w http.ResponseWriter, err error) {
	sendError(w, http.StatusGatewayTimeout, err)
}

// NotModified sends an error with status 304 Not Modified
func NotModified(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotModified, err)
}

// ExpectationFailed sends an error with status 417 Expectation Failed
func ExpectationFailed(w http.ResponseWriter, err error) {
	sendError(w, http.StatusExpectationFailed, err)
}

// BadData sends an error with status 422 Unprocessable Entity
func BadData(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnprocessableEntity, err)
}

// BadRequest sends an error with status 400 Bad Request
func BadRequest(w http.ResponseWriter, err error) {
	sendError(w, http.StatusBadRequest, err)
}

// Conflict sends an error with status 409 Conflict
func Conflict(w http.ResponseWriter, err error) {
	sendError(w, http.StatusConflict, err)
}

// Forbidden sends an error with status 403 Forbidden
func Forbidden(w http.ResponseWriter, err error) {
	sendError(w, http.StatusForbidden, err)
}

// InternalServerError sends an error with status 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, err error) {
	sendError(w, http.StatusInternalServerError, err)
}

// NotFound sends an error with status 404 Not Found
func NotFound(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotFound, err)
}

// Unauthorized sends an error with status 401 Unauthorized
func Unauthorized(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnauthorized, err)
}
