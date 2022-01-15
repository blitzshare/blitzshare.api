package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestP2pRegistry(t *testing.T) {
	url := os.Getenv("APP_URL")
	ack := PostPeerRegistry(t, &url)
	assert.NotNil(t, ack.AckId)
	assert.True(t, len(ack.AckId) > 5)
	// since we're in event driven, we need to wasit some time for the request to be processed
	time.Sleep(time.Second * 2)
	addr := GetPeerRegistry(t, &url)
	assert.NotNil(t, addr.MultiAddr)
	assert.NotEmpty(t, addr.MultiAddr)

	nodeConfig := GetBootstrapNode(t, &url)
	fmt.Println(nodeConfig.NodeId)
	assert.NotNil(t, nodeConfig.NodeId)
	assert.Equal(t, nodeConfig.Port, 63785)
	// deregister
	deleteAck := DeletePeerRegistry(t, url, ack.Token)
	assert.NotNil(t, deleteAck.AckId)
	assert.True(t, len(deleteAck.AckId) > 5)
}
