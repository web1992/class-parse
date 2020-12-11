package parse

type Parse interface {
	parseFile(fileName string) error
	Bytes() []byte
	Name() string
}
