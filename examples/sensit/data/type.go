package data

import (
	"strconv"

	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/errs"
)

type Type struct {
	Code uint8
	Text string
}

var (
	RegularType = &Type{
		Code: 0,
		Text: "Regular",
	}
	ButtonCallType = &Type{
		Code: 1,
		Text: "ButtonCall",
	}
	AlertType = &Type{
		Code: 2,
		Text: "Alert",
	}
	NewModeType = &Type{
		Code: 3,
		Text: "NewMode",
	}
)

func GetType(b byte) (*Type, *errs.SensitError) {
	typeCode := (b & 0x60) >> 5
	switch typeCode {
	case RegularType.Code:
		return RegularType, nil
	case ButtonCallType.Code:
		return ButtonCallType, nil
	case AlertType.Code:
		return AlertType, nil
	case NewModeType.Code:
		return NewModeType, nil
	default:
		return nil, errs.NewSensitError("invalid type: "+strconv.Itoa(int(typeCode)), errs.UnknownTypeErrorCode)
	}
}
