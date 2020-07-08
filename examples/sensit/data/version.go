package data

type Version struct {
	Major uint8
	Minor uint8
}

func NewVersion(b byte) *Version {
	return &Version{
		Major: (b & 0xf0) >> 4,
		Minor: b & 0x0f,
	}
}
