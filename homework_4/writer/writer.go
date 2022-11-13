package writer

import "fmt"

var CustomWriter *customWriter

func init() {
	CustomWriter = NewCustomWriter()
}

type customWriter struct{}

func NewCustomWriter() *customWriter {
	return &customWriter{}
}

func (c *customWriter) Write(p []byte) (n int, err error) {
	fmt.Println("[CUSTOM WRITER]", string(p))
	return len(p), nil
}
