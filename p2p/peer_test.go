package p2p

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mendozg/not-only-mining-pool/config"
)

func TestNewPeer(t *testing.T) {
	var options config.P2POptions
	json.Unmarshal([]byte(`
{
    "host": "0.0.0.0",
    "port": 19335,
    "magic": "fdd2c8f1",
    "disableTransactions": true
}
`), &options)
	peer := NewPeer(70015, &options)
	peer.Init()

	c := time.After(time.Minute)
	for {
		select {
		case <-c:
			return
		}
	}
}
