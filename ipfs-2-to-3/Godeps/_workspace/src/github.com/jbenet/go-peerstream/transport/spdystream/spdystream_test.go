package peerstream_spdystream

import (
	"testing"

	psttest "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-peerstream/transport/test"
)

func TestSpdyStreamTransport(t *testing.T) {
	t.Skip("spdystream is known to be broken")
	psttest.SubtestAll(t, Transport)
}
