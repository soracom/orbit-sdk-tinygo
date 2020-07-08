package results

type JSONSerializable interface {
	SerializeJSON() ([]byte, error)
}
