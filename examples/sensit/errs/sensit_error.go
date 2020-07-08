package errs

import sdk "github.com/soracom/orbit-sdk-tinygo"

type SensitErrorCode sdk.ErrorCode

const (
	OkErrorCode                      SensitErrorCode = 0
	InsufficientDataErrorCode                        = -1
	InvalidBase64ErrorCode                           = -2
	UnknownModeErrorCode                             = -10
	UnknownTimeframeErrorCode                        = -11
	UnknownTypeErrorCode                             = -12
	UnableToSerializeToJSONErrorCode                 = -100
	UnableToGetBufferErrorCode                       = -101
)

type SensitError struct {
	message   string
	errorCode SensitErrorCode
}

func NewSensitError(msg string, errorCode SensitErrorCode) *SensitError {
	return &SensitError{
		message:   msg,
		errorCode: errorCode,
	}
}

func (e *SensitError) Error() string {
	return e.message
}

func (e *SensitError) ErrorCode() SensitErrorCode {
	return e.errorCode
}
