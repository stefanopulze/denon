package packager

type Packager interface {
	Encode(string) ([]byte, error)

	Verify([]byte, []byte) error

	Decode([]byte) (string, error)
}
