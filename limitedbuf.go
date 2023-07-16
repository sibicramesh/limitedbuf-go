package limitedbuf

import (
	"bytes"
	"errors"
	"io"
)

var _ io.Writer = &LimitedBuf{}
var _ Buffer = &LimitedBuf{}
var _ Buffer = &bytes.Buffer{}

// ErrWriteExceedsBufCap is returned when write exceeds buffer capacity.
var ErrWriteExceedsBufCap = errors.New("exceeds buffer capacty")

// Buffer implements core methods to manipulate a bytes slice.
type Buffer interface {
	Write(p []byte) (int, error)
	Bytes() []byte
	Len() int
}

// LimitedBuf implements io.Writer interface.
type LimitedBuf struct {
	buf *bytes.Buffer
	max int
}

// NewBuffer creates a new buffer which implements io.Writer interface.
// The bytes slice passed should be allocated with
// Len: 0
// Cap: default
// You cannot pass nil and expect the buffer to created.
// The max argument specifies the max number of bytes to allow in the buffer.
func NewBuffer(buf []byte, max int) *LimitedBuf {
	return &LimitedBuf{
		buf: bytes.NewBuffer(buf),
		max: max,
	}
}

// Write first checks if there is enough capacity in the buffer to write p,
// If the validation fails, Write method returns 0 and ErrWriteExceedsBufCap.
func (b *LimitedBuf) Write(p []byte) (int, error) {

	if len(p)+b.buf.Len() > b.max {
		return 0, ErrWriteExceedsBufCap
	}

	return b.buf.Write(p)
}

// Bytes is wrapper which calls bytes.Bytes()
func (b *LimitedBuf) Bytes() []byte {
	return b.buf.Bytes()
}

// Len is wrapper which calls bytes.Len()
func (b *LimitedBuf) Len() int {
	return b.buf.Len()
}
