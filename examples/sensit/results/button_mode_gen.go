// Code generated by "json-ice --type=ButtonMode"; DO NOT EDIT.

package results

import "github.com/moznion/go-json-ice/serializer"

func MarshalButtonModeAsJSON(s *ButtonMode) ([]byte, error) {
	buff := make([]byte, 1, 366)
	buff[0] = '{'
	buff = append(buff, "\"mode\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Mode))
	buff = append(buff, ',')
	buff = append(buff, "\"mode_text\":"...)
	buff = serializer.AppendSerializedString(buff, s.ModeText)
	buff = append(buff, ',')
	buff = append(buff, "\"timeframe\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Timeframe))
	buff = append(buff, ',')
	buff = append(buff, "\"timeframe_text\":"...)
	buff = serializer.AppendSerializedString(buff, s.TimeframeText)
	buff = append(buff, ',')
	buff = append(buff, "\"type\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.Type))
	buff = append(buff, ',')
	buff = append(buff, "\"type_text\":"...)
	buff = serializer.AppendSerializedString(buff, s.TypeText)
	buff = append(buff, ',')
	buff = append(buff, "\"battery\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.Battery))
	buff = append(buff, ',')
	buff = append(buff, "\"temp_c_low_precision\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.TempCLowPrecision))
	buff = append(buff, ',')
	buff = append(buff, "\"temp_c\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.TempC))
	buff = append(buff, ',')
	buff = append(buff, "\"temp_f_low_precision\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.TempFLowPrecision))
	buff = append(buff, ',')
	buff = append(buff, "\"temp_f\":"...)
	buff = serializer.AppendSerializedFloat(buff, float64(s.TempF))
	buff = append(buff, ',')
	buff = append(buff, "\"reed_switch_state\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.ReedSwitchState))
	buff = append(buff, ',')
	buff = append(buff, "\"major_version\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.MajorVersion))
	buff = append(buff, ',')
	buff = append(buff, "\"minor_version\":"...)
	buff = serializer.AppendSerializedUint(buff, uint64(s.MinorVersion))
	buff = append(buff, ',')
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}