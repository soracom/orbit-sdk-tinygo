package errs

import sdk "github.com/soracom/orbit-sdk-tinygo"

type LTEMButtonErrorCode sdk.ErrorCode

const (
	OkErrorCode   LTEMButtonErrorCode = 0
	ExecErrorCode                     = -1
)
