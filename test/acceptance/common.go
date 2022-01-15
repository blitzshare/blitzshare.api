package acceptance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"blitzshare.api/app/model"
	"github.com/stretchr/testify/assert"
)

func DeletePeerRegistry(t *testing.T, baseUrl, token string) *model.AckResponse {
	url := fmt.Sprintf("%s/p2p/registry/%s/%s", baseUrl, Otp, token)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer response.Body.Close()
	ack := model.AckResponse{}
	b, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)
	json.Unmarshal(b, &ack)
	return &ack
}

func PostPeerRegistry(t *testing.T, url *string) *model.PeerRegistryAckResponse {
	body, _ := json.Marshal(model.P2pPeerRegistryReq{
		MultiAddr: MultiAddr,
		Otp:       Otp,
		Mode:      "chat",
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
