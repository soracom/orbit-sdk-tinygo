package parsers

import (
	"strconv"

	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/errs"
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/results"
)

type Parser interface {
	Parse(input []byte) (results.JSONSerializable, *errs.SensitError)
}

func GetParser(b byte) (Parser, *errs.SensitError) {
	switch b & 0x07 {
	case ButtonModeCode:
		return &ButtonModeParser{}, nil
	default:
		return nil, errs.NewSensitError("unknown mode: "+strconv.Itoa(int(b)), errs.UnknownModeErrorCode)
	}
}
