package acceptance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"blitzshare.api/app/model"
	"github.com/stretchr/testify/assert"
)

func PostPeerRegistry(t *testing.T, url *string) *model.PeerRegistryAckResponse {
	body, _ := json.Marshal(model.P2pPeerRegistryCmd{
		MultiAddr: MultiAddr,
		Otp:       Otp,
	})
	serverUrl := fmt.Sprintf("%s/p2p/registry", *url)
	resp, err := http.Post(serverUrl, "application/json", bytes.NewReader(body))

	ack := model.PeerRegistryAckResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &ack)
	return &ack
}

func GetPeerRegistry(t *testing.T, url *string) *model.MultiAddrResponse {
	serverUrl := fmt.Sprintf("%s/p2p/registry/%s", *url, Otp)
	resp, err := http.Get(serverUrl)
	addr := model.MultiAddrResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println("@@@", string(b))
	assert.Nil(t, err)
	json.Unmarshal(b, &addr)
	return &addr
}

func GetBootstrapNode(t *testing.T, url *string) *model.NodeConfigRespone {
	serverUrl := fmt.Sprintf("%s/p2p/bootstrap-node", *url)
	resp, err := http.Get(serverUrl)
	addr := model.NodeConfigRespone{}
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &addr)
	return &addr
}
