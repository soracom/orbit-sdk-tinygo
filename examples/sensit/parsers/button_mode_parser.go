package parsers

import (
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/data"
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/errs"
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/results"
)

const (
	ButtonModeCode uint8  = 0
	ButtonModeText string = "Button"
)

type ButtonModeParser struct {
}

func (p *ButtonModeParser) Parse(input []byte) (results.JSONSerializable, *errs.SensitError) {
	if len(input) < 4 {
		return nil, errs.NewSensitError("insufficient input data", errs.InsufficientDataErrorCode)
	}
	b0 := input[0]
	b1 := input[1]
	b2 := input[2]
	b3 := input[3]

	timeframe, sensitError := data.GetTimeframe(b0)
	if sensitError != nil {
		return nil, sensitError
	}

	typ, sensitError := data.GetType(b0)
	if sensitError != nil {
		return nil, sensitError
	}

	battery := data.GetBattery(b0, b1)
	temperature := data.NewTemperature(b1, b2)
	reedSwitchState := data.GetReedSwitchState(b2)
	version := data.NewVersion(b3)

	return &results.ButtonMode{
		Mode:              ButtonModeCode,
		ModeText:          ButtonModeText,
		Timeframe:         timeframe.Code,
		TimeframeText:     timeframe.Text,
		Type:              typ.Code,
		TypeText:          typ.Text,
		Battery:           battery,
		TempCLowPrecision: temperature.GetTempCLowPrecision(),
		TempC:             temperature.GetTempC(),
		TempFLowPrecision: temperature.GetTempFLowPrecision(),
		TempF:             temperature.GetTempF(),
		ReedSwitchState:   reedSwitchState,
		MajorVersion:      version.Major,
		MinorVersion:      version.Minor,
	}, nil
}
