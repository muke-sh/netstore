package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTrasport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr: ":3000",
	}

	tr := NewTCPTransport(opts)

	assert.Nil(t, tr.ListenAndAccept())
}
