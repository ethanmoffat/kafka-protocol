package protocol

type Message interface {
	Marshal() []byte
	Unmarshal([]byte) error

	ApiKey() ApiKey
	Version() int
	CorrelationId() int
}
