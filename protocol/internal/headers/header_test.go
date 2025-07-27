package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaders(t *testing.T) {
	// Test: Valid single header
	h := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := h.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	assert.Equal(t, "localhost:42069", h.Get("Host"))
	assert.Equal(t, 25, n) // count the registered nurse too
	assert.True(t, done)

	// Test: Valid single header
	h = NewHeaders()
	data = []byte("Host: localhost:42069\r\n HEy:    MOhajana   \r\n\r\n")
	n, done, err = h.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	assert.Equal(t, "localhost:42069", h.Get("Host"))
	assert.Equal(t, "MOhajana", h.Get("HEy"))
	assert.Equal(t, 47, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	h = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = h.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid spacing header
	h = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = h.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Valid headers sharing common header
	h = NewHeaders()
	data = []byte("User: I'm fine\r\n user: I'm fine too   \r\n\r\n")
	n, done, err = h.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	assert.Equal(t, "I'm fine, I'm fine too", h.Get("user"))
	assert.Equal(t, 42, n)
	assert.True(t, done)
}
