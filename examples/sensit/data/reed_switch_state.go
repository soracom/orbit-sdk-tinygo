package data

func GetReedSwitchState(b byte) uint8 {
	return (b & 0x40) >> 6
}
