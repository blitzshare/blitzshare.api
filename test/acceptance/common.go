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

func PostPeerRegistry(t *testing.T) *model.PeerRegistryAckResponse {
	body, _ := json.Marshal(model.P2pPeerRegistryCmd{
		MultiAddr: MultiAddr,
		Otp:       "secret-pass",
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
