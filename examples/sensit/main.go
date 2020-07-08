package main

import (
	sdk "github.com/soracom/orbit-sdk-tinygo"
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/errs"
	"github.com/soracom/orbit-sdk-tinygo/examples/sensit/parsers"
)

func main() {
}

//export uplink
func uplink() sdk.ErrorCode {
	inputBuffer, err := sdk.GetInputBuffer()
	if err != nil {
		sdk.Log(err.Error())
		return errs.UnableToGetBufferErrorCode
	}

	if len(inputBuffer) < 0 {
		sdk.Log("insufficient input buffer; cannot get mode information")
		return errs.InsufficientDataErrorCode
	}
	parser, serr := parsers.GetParser(inputBuffer[0])
	if serr != nil {
		sdk.Log(serr.Error())
		return sdk.ErrorCode(serr.ErrorCode())
	}

	parsed, serr := parser.Parse(inputBuffer)
	if serr != nil {
		sdk.Log(serr.Error())
		return sdk.ErrorCode(serr.ErrorCode())
	}

	serialized, err := parsed.SerializeJSON()
	if err != nil {
		sdk.Log("error while serializing button mode to json: " + err.Error())
		return errs.UnableToSerializeToJSONErrorCode
	}

	sdk.SetOutputJSON(string(serialized))

	return sdk.ErrorCode(errs.OkErrorCode)
}
