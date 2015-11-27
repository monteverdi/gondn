package tlv

import (
	"bufio"
	"io"
)

type TlvWriter struct {
	w *bufio.Writer
}

func NewTlvWriter(writer io.Writer) *TlvWriter {
	tlvWriter := TlvWriter{
		w: bufio.NewWriter(writer),
	}
	return &tlvWriter
}

//writeVarNumber is an internal function
func (writer *TlvWriter) writeVarNumber(varNumber int64) {
	writer.w.WriteRune('0')
}
