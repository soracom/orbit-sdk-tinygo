package main

import (
	"github.com/moznion/jsonparser"
	sdk "github.com/soracom/orbit-sdk-tinygo"
	"github.com/soracom/orbit-sdk-tinygo/examples/lte-m-button/data"
	"github.com/soracom/orbit-sdk-tinygo/examples/lte-m-button/errs"
)

func main() {
}

//export uplink
func uplink() sdk.ErrorCode {
	inputBuffer, err := sdk.GetInputBuffer()
	if err != nil {
		sdk.Log(err.Error())
		return sdk.ErrorCode(errs.ExecErrorCode)
	}

	output, err := convertInputToOutput(inputBuffer)
	if err != nil {
		sdk.Log(err.Error())
		return sdk.ErrorCode(errs.ExecErrorCode)
	}

	serializedOutput, err := output.SerializeJSON()
	if err != nil {
		sdk.Log(err.Error())
		return sdk.ErrorCode(errs.ExecErrorCode)
	}

	sdk.SetOutputJSON(string(serializedOutput))

	return sdk.ErrorCode(errs.OkErrorCode)
}

func convertInputToOutput(input []byte) (*data.Output, error) {
	clickType, err := jsonparser.GetInt(input, "clickType")
	if err != nil {
		return nil, err
	}

	clickTypeName, err := jsonparser.GetString(input, "clickTypeName")
	if err != nil {
		return nil, err
	}

	batteryLevel, err := jsonparser.GetFloat(input, "batteryLevel")
	if err != nil {
		return nil, err
	}

	binaryParserEnabled, err := jsonparser.GetBoolean(input, "binaryParserEnabled")
	if err != nil {
		return nil, err
	}

	imsi, err := sdk.GetSourceValue("resourceId")
	if err != nil {
		return nil, err
	}

	name, err := sdk.GetTagValue("name")
	if err != nil {
		return nil, err
	}

	location, _ := sdk.GetLocation() // if location doesn't exist (i.e. location is nil), it returns an error as the second return value

	timestamp := sdk.GetTimestamp()

	return &data.Output{
		ClickType:           int32(clickType),
		ClickTypeName:       clickTypeName,
		BatteryLevel:        batteryLevel,
		BinaryParserEnabled: binaryParserEnabled,
		Imsi:                string(imsi),
		Name:                string(name),
		Location:            location,
		Timestamp:           timestamp,
	}, nil
}
