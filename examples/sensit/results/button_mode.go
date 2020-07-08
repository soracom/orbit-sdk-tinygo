package results

//go:generate json-ice --type=ButtonMode
type ButtonMode struct {
	Mode              uint8   `json:"mode"`
	ModeText          string  `json:"mode_text"`
	Timeframe         uint8   `json:"timeframe"`
	TimeframeText     string  `json:"timeframe_text"`
	Type              uint8   `json:"type"`
	TypeText          string  `json:"type_text"`
	Battery           float32 `json:"battery"`
	TempCLowPrecision float32 `json:"temp_c_low_precision"`
	TempC             float32 `json:"temp_c"`
	TempFLowPrecision float32 `json:"temp_f_low_precision"`
	TempF             float32 `json:"temp_f"`
	ReedSwitchState   uint8   `json:"reed_switch_state"`
	MajorVersion      uint8   `json:"major_version"`
	MinorVersion      uint8   `json:"minor_version"`
}

func (s *ButtonMode) SerializeJSON() ([]byte, error) {
	return MarshalButtonModeAsJSON(s)
}
