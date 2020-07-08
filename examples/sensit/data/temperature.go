package data

type Temperature struct {
	msb byte
	lsb byte
}

func NewTemperature(b1, b2 byte) *Temperature {
	return &Temperature{
		msb: (b1 & 0xf0) >> 4,
		lsb: b2 & 0x3f,
	}
}

func (t *Temperature) GetTempC() float32 {
	return getTempC(t.msb, t.lsb)
}

func (t *Temperature) GetTempF() float32 {
	return (t.GetTempC() * 1.8) + 32.0
}

func (t *Temperature) GetTempCLowPrecision() float32 {
	return getTempC(t.msb, 0)
}

func (t *Temperature) GetTempFLowPrecision() float32 {
	return (t.GetTempCLowPrecision() * 1.8) + 32.0
}

func getTempC(msb, lsb byte) float32 {
	return (((float32(msb) * 64) + float32(lsb)) - 200) / 8.0
}
