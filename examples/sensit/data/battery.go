package data

func GetBattery(b1, b2 byte) float32 {
	msb := float32((b1 & 0x80) >> 7)
	lsb := float32(b2 & 0x0f)

	return (((msb * 16) + lsb) + 54) / 20.0
}
