package data

import (
	"strconv"

	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/errs"
)

type Timeframe struct {
	Code uint8
	Text string
}

var (
	TenMinutesTimeframe = &Timeframe{
		Code: 0,
		Text: "TenMinutes",
	}
	OneHourTimeframe = &Timeframe{
		Code: 1,
		Text: "OneHour",
	}
	SixHoursTimeframe = &Timeframe{
		Code: 2,
		Text: "SixHours",
	}
	TwentyFourHoursTimeframe = &Timeframe{
		Code: 3,
		Text: "TwentyFourHours",
	}
)

func GetTimeframe(b byte) (*Timeframe, *errs.SensitError) {
	timeframeCode := (b & 0x18) >> 3
	switch timeframeCode {
	case TenMinutesTimeframe.Code:
		return TenMinutesTimeframe, nil
	case OneHourTimeframe.Code:
		return OneHourTimeframe, nil
	case SixHoursTimeframe.Code:
		return SixHoursTimeframe, nil
	case TwentyFourHoursTimeframe.Code:
		return TwentyFourHoursTimeframe, nil
	default:
		return nil, errs.NewSensitError("invalid timeframe: "+strconv.Itoa(int(timeframeCode)), errs.UnknownTimeframeErrorCode)
	}
}
