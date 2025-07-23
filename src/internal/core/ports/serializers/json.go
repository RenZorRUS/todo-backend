package serializers

type JSONSerializer interface {
	Marshal(value any) ([]byte, error)
	Unmarshal(data []byte, value any) error
}
