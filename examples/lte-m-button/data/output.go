package data

import sdk "github.com/soracom/orbit-sdk-tinygo"

//go:generate json-ice --type=Output
type Output struct {
	ClickType           int32         `json:"clickType"`
	ClickTypeName       string        `json:"clickTypeName"`
	BatteryLevel        float64       `json:"batteryLevel"`
	BinaryParserEnabled bool          `json:"binaryParserEnabled"`
	Imsi                string        `json:"imsi"`
	Name                string        `json:"name"`
	Location            *sdk.Location `json:"location"`
	Timestamp           int64         `json:"timestamp"`
	UserData            string        `json:"userData"`
}

func (o *Output) SerializeJSON() ([]byte, error) {
	return MarshalOutputAsJSON(o)
}
