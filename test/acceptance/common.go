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

//curl -XPOST a4d8a61f98aae42179b2567972dfe9d1-1191192214.eu-west-2.elb.amazonaws.com/p2p/registry -d '{ "multiAddr": "x", "otp": "y0}'
func PostPeerRegistry(t *testing.T) *model.PeerRegistryAckResponse {
	body, _ := json.Marshal(model.P2pPeerRegistryCmd{
		MultiAddr: MultiAddr,
		Otp:       Otp,
	})
	serverUrl := fmt.Sprintf("%s/p2p/registry", Url)
	resp, err := http.Post(serverUrl, "application/json", bytes.NewReader(body))

	ack := model.PeerRegistryAckResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &ack)
	return &ack
}

func GetPeerRegistry(t *testing.T) *model.MultiAddrResponse {
	serverUrl := fmt.Sprintf("%s/p2p/registry/%s", Url, Otp)
	resp, err := http.Get(serverUrl)
	addr := model.MultiAddrResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &addr)
	return &addr
}

func GetBootstrapNode(t *testing.T) *model.NodeConfigRespone {
	serverUrl := fmt.Sprintf("%s/p2p/bootstrap-node", Url)
	resp, err := http.Get(serverUrl)
	addr := model.NodeConfigRespone{}
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &addr)
	return &addr
}
