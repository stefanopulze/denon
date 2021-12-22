package transport

type Transport interface {
	Connect() error

	Send([]byte) ([]byte, error)

	Close() error
}
