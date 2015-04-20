package peerstream_yamux

import (
	"testing"

	psttest "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-peerstream/transport/test"
)

func TestYamuxTransport(t *testing.T) {
	psttest.SubtestAll(t, DefaultTransport)
}
