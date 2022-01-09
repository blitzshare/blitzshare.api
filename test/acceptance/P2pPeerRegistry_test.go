package acceptance

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestP2pRegistry(t *testing.T) {
	ack := PostPeerRegistry(t)
	assert.NotNil(t, ack.AckId)
	assert.True(t, len(ack.AckId) > 5)
	// since we're in event driven, we need to wasit some time for the request to be processed
	time.Sleep(time.Second * 2)
	addr := GetPeerRegistry(t)
	fmt.Println(addr)
	assert.NotNil(t, addr.MultiAddr)
	assert.NotEmpty(t, addr.MultiAddr)

	nodeConfig := GetBootstrapNode(t)
	fmt.Println(nodeConfig.NodeId)
	assert.NotNil(t, nodeConfig.NodeId)
	assert.Equal(t, nodeConfig.Port, 63785)
}
