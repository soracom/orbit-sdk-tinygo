package sdk

import (
	"errors"
)

var (
	ErrNoInputBuffer            = errors.New("no input buffer")
	ErrInvalidInputBufferLength = errors.New("expected input buffer length and actual buffer length are different")
	ErrNoTagValue               = errors.New("no tag value")
	ErrInvalidTagValueLength    = errors.New("expected tag value buffer length and actual buffer length are different")
	ErrNoSourceValue            = errors.New("no source value")
	ErrInvalidSourceValueLength = errors.New("expected source value buffer length and actual buffer length are different")
	ErrNoLocationInformation    = errors.New("there is no location information")
)

// ErrorCode represents the error code for Soralet.
type ErrorCode int32

// Location is a struct that contains the longitude and latitude information.
//
//go:generate json-ice --type=Location
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (l *Location) MarshalJSON() ([]byte, error) {
	return MarshalLocationAsJSON(l)
}

// Following functions are provided by orbit runtime. {{{

//go:wasm-module env
//export orbit_log
func orbitLog(string)

//go:wasm-module env
//export orbit_get_input_buffer
func orbitGetInputBuffer(*byte, int32) int32

//go:wasm-module env
//export orbit_get_input_buffer_len
func orbitGetInputBufferLen() int32

//go:wasm-module env
//export orbit_set_output
func orbitSetOutput(string)

//go:wasm-module env
//export orbit_set_output_content_type
func orbitSetOutputContentType(string)

//go:wasm-module env
//export orbit_get_tag_value
func orbitGetTagValue(string, *byte, int32) int32

//go:wasm-module env
//export orbit_get_tag_value_len
func orbitGetTagValueLen(string) int32

//go:wasm-module env
//export orbit_get_source_value
func orbitGetSourceValue(string, *byte, int32) int32

//go:wasm-module env
//export orbit_get_source_value_len
func orbitGetSourceValueLen(string) int32

//go:wasm-module env
//export orbit_has_location
func orbitHasLocation() int32

//go:wasm-module env
//export orbit_get_location_lat
func orbitGetLocationLat() float64

//go:wasm-module env
//export orbit_get_location_lon
func orbitGetLocationLon() float64

//go:wasm-module env
//export orbit_get_timestamp
func orbitGetTimestamp() int64

//go:wasm-module env
//export orbit_get_user_data
func orbitGetUserdata(*byte, int32) int32

//go:wasm-module env
//export orbit_get_userdata_len
func orbitGetUserdataLen() int32

//go:wasm-module env
//export orbit_get_original_request
func orbitGetOriginalRequest(*byte, int32) int32

//go:wasm-module env
//export orbit_get_original_request_len
func orbitGetOriginalRequestLen() int32

// }}}

// Log outputs a log entry.
func Log(msg string) {
	orbitLog(msg)
}

// GetInputBuffer retrieves an input value from the orbit backend.
func GetInputBuffer() ([]byte, error) {
	bufferLen := orbitGetInputBufferLen()
	if bufferLen <= 0 {
		return nil, ErrNoInputBuffer
	}

	buff := make([]byte, bufferLen, bufferLen)
	actualLen := orbitGetInputBuffer(&buff[0], bufferLen)
	if bufferLen != actualLen {
		return nil, ErrInvalidInputBufferLength
	}

	return buff, nil
}

// GetTagValue retrieves tag value according to the given name.
func GetTagValue(name string) ([]byte, error) {
	bufferLen := orbitGetTagValueLen(name)
	if bufferLen <= 0 {
		return nil, ErrNoTagValue
	}

	buff := make([]byte, bufferLen, bufferLen)
	actualLen := orbitGetTagValue(name, &buff[0], bufferLen)
	if bufferLen != actualLen {
		return nil, ErrInvalidTagValueLength
	}

	return buff, nil
}

// GetSourceValue retrieves source value according to the given name.
func GetSourceValue(name string) ([]byte, error) {
	bufferLen := orbitGetSourceValueLen(name)
	if bufferLen <= 0 {
		return nil, ErrNoSourceValue
	}
	buff := make([]byte, bufferLen, bufferLen)
	actualLen := orbitGetSourceValue(name, &buff[0], bufferLen)
	if bufferLen != actualLen {
		return nil, ErrInvalidSourceValueLength
	}

	return buff, nil
}

func GetUserData() ([]byte, error) {
	bufferLen := orbitGetUserdataLen()
	if bufferLen <= 0 {
		return make([]byte, 0), nil
	}
	buff := make([]byte, bufferLen, bufferLen)
	actualLen := orbitGetUserdata(&buff[0], bufferLen)
	if bufferLen != actualLen {
		return nil, ErrInvalidSourceValueLength
	}

	return buff, nil
}

func GetOriginalRequest() ([]byte, error) {
	bufferLen := orbitGetOriginalRequestLen()
	if bufferLen <= 0 {
		return make([]byte, 0), nil
	}
	buff := make([]byte, bufferLen, bufferLen)
	actualLen := orbitGetOriginalRequest(&buff[0], bufferLen)
	if bufferLen != actualLen {
		return nil, ErrInvalidSourceValueLength
	}

	return buff, nil
}

// GetLocation retrieves location information if that exists.
// If the location information doesn't exist, this function returns a nil struct and "no-location" error.
func GetLocation() (*Location, error) {
	if orbitHasLocation() == 0 {
		return nil, ErrNoLocationInformation
	}

	return &Location{
		Lat: orbitGetLocationLat(),
		Lon: orbitGetLocationLon(),
	}, nil
}

// GetTimestamp returns the timestamp.
func GetTimestamp() int64 {
	return orbitGetTimestamp()
}

// SetOutputJSON sets a JSON string as the output for the orbit backend.
func SetOutputJSON(out string) {
	contentType := "application/json"
	orbitSetOutputContentType(contentType)
	orbitSetOutput(out)
}
