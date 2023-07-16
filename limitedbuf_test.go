package limitedbuf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LimitedBuf(t *testing.T) {

	b := make([]byte, 0)
	buf := NewBuffer(b, 5)
	require.NotNil(t, buf)

	n, err := buf.Write([]byte{1, 2, 3, 4})
	require.Nil(t, err)
	require.Equal(t, 4, n)
	require.Equal(t, 4, buf.Len())

	b = make([]byte, 0)
	buf = NewBuffer(b, 2)

	n, err = buf.Write([]byte{1, 2, 3, 4})
	require.Equal(t, ErrWriteExceedsBufCap, err)
	require.Equal(t, 0, n)
	require.Equal(t, 0, buf.Len())

	b = make([]byte, 0)
	buf = NewBuffer(b, 4)

	n, err = buf.Write([]byte{1, 2})
	require.Nil(t, err)
	require.Equal(t, 2, n)
	require.Equal(t, 2, buf.Len())

	n, err = buf.Write([]byte{3, 4})
	require.Nil(t, err)
	require.Equal(t, 2, n)
	require.Equal(t, 4, buf.Len())

	n, err = buf.Write([]byte{6, 7})
	require.Equal(t, ErrWriteExceedsBufCap, err)
	require.Equal(t, 0, n)
	require.Equal(t, 4, buf.Len())
}
