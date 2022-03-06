package test

import (
	"blitzshare.api/app/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cucumber/godog"
	_ "github.com/lib/pq"
)

const (
	MultiAddr = "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"
	Otp       = "test-reimpose-verminosis-acidulate"
	ApiKey    = "blitzshare-client-XCVsdfsdfSDFxcvWErsd3"
)

var client = &http.Client{}

var baseUrl string

type healthCheckStatusCodeKey struct{}
type postPeerRegistryStatusCode struct{}

func getHealthCheck(ctx context.Context) context.Context {
	url := fmt.Sprintf("%s/test", os.Getenv("API_URL"))
	r, _ := http.Get(url)
	return context.WithValue(ctx, healthCheckStatusCodeKey{}, r.StatusCode)
}

func assertHealthCheckResponseStatusIsOk(ctx context.Context) error {
	statusCode := ctx.Value(healthCheckStatusCodeKey{}).(int)
	if statusCode == http.StatusOK {
		return nil
	}
	return errors.New("status code is not OK")
}

func assertStatusCodeIsUnauthorized(ctx context.Context) error {
	statusCode := ctx.Value(postPeerRegistryStatusCode{}).(int)
	if statusCode == http.StatusUnauthorized {
		return nil
	}
	return errors.New("status code is not OK")
}

func postPeerRegistry(ctx context.Context, authKey bool) context.Context {
	body, _ := json.Marshal(model.P2pPeerRegistryReq{
		MultiAddr: model.MultiAddr{
			MultiAddr: MultiAddr,
		},
		Otp: model.Otp{
			Otp: Otp,
		},
		Mode: "chat",
	})
	serverUrl := fmt.Sprintf("%s/p2p/registry", baseUrl)

	req, _ := http.NewRequest("POST", serverUrl, bytes.NewReader(body))
	if authKey {
		req.Header.Set("x-api-key", ApiKey)
	}

	resp, _ := client.Do(req)

	ack := model.PeerRegistryAckResponse{}
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &ack)

	return context.WithValue(
		context.WithValue(ctx, postPeerRegistryStatusCode{}, resp.StatusCode),
		model.PeerRegistryAckResponse{}, ack)
}

func getPeerInfoViaOTP(ctx context.Context) context.Context {
	serverUrl := fmt.Sprintf("%s/p2p/registry/%s", baseUrl, Otp)
	req, _ := http.NewRequest("GET", serverUrl, nil)
	req.Header.Set("x-api-key", ApiKey)
	resp, _ := client.Do(req)
	otpRegistry := model.P2pPeerRegistryResponse{}
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &otpRegistry)
	return context.WithValue(ctx, model.P2pPeerRegistryResponse{}, otpRegistry)
}
func getBootstrapNodeConfig(ctx context.Context) context.Context {
	serverUrl := fmt.Sprintf("%s/p2p/bootstrap-node", baseUrl)
	req, _ := http.NewRequest("GET", serverUrl, nil)
	req.Header.Set("x-api-key", ApiKey)
	resp, _ := client.Do(req)
	nodeConfig := model.NodeConfigRespone{}
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &nodeConfig)
	return context.WithValue(ctx, model.NodeConfigRespone{}, nodeConfig)
}
func validateTestContext(ctx context.Context) error {
	nodeConfig := ctx.Value(model.NodeConfigRespone{}).(model.NodeConfigRespone)
	//fmt.Println("nodeConfig", nodeConfig)
	if nodeConfig.NodeId == "" || nodeConfig.Port == 0 {
		return errors.New("invalid node config")
	}
	otpRegistry := ctx.Value(model.P2pPeerRegistryResponse{}).(model.P2pPeerRegistryResponse)
	//fmt.Println("otpRegistry", otpRegistry)
	if otpRegistry.MultiAddr.MultiAddr == "" || otpRegistry.Mode.Mode == "" || otpRegistry.Otp.Otp == "" {
		return errors.New("invalid otpRegistry")
	}
	registry := ctx.Value(model.PeerRegistryAckResponse{}).(model.PeerRegistryAckResponse)
	// fmt.Println("registry", registry)
	if registry.AckId == "" || otpRegistry.Otp.Otp == "" {
		return errors.New("invalid registry")
	}
	return nil
}

func deleteUserRegistration(ctx context.Context) error {
	registry := ctx.Value(model.PeerRegistryAckResponse{}).(model.PeerRegistryAckResponse)
	url := fmt.Sprintf("%s/p2p/registry/%s/%s", baseUrl, Otp, registry.Token)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("x-api-key", ApiKey)
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if response.StatusCode != http.StatusAccepted {
		return err
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	baseUrl = os.Getenv("API_URL")
	ctx.Step(`^http GET health check request executed$`, getHealthCheck)
	ctx.Step(`^http response is OK$`, assertHealthCheckResponseStatusIsOk)
	ctx.Step(`^User registers via OTP$`, func(ctx context.Context) context.Context {
		return postPeerRegistry(ctx, true)
	})
	ctx.Step(`^Another User obtains registred user information via OTP$`, getPeerInfoViaOTP)
	ctx.Step(`^User get bootstrap node config$`, getBootstrapNodeConfig)
	ctx.Step(`^Connection between users can be etablished$`, validateTestContext)
	ctx.Step(`^User can deregister OTP via obtained Token$`, deleteUserRegistration)
	ctx.Step(`^User registers via OTP without auth header$`, func(ctx context.Context) context.Context {
		return postPeerRegistry(ctx, false)
	})
	ctx.Step(`^User request is unauthorized`, assertStatusCodeIsUnauthorized)
}
