package endpoints_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server"
	"blitzshare.api/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Register P2p Endpoint", func() {
	const (
		OTP       = "gelandelaufer-astromancer-scurvyweed-sayability"
		MultiAddr = "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"
		AckId     = "12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"
	)
	var deps *dependencies.Dependencies
	BeforeEach(func() {
		server := config.Server{Port: 323}
		settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
		emit := &mocks.EventEmit{}
		emit.On("EmitP2pPeerRegisterCmd",
			mock.MatchedBy(func(input interface{}) bool {
				return true
			}),
			mock.MatchedBy(func(input interface{}) bool {
				return true
			}),
			mock.MatchedBy(func(input interface{}) bool {
				return true
			}),
		).Return(AckId, nil)
		config := config.Config{
			Server:   server,
			Settings: settings,
		}
		rndMock := &mocks.Rnd{}
		r := "random-string"
		rndMock.On("GenerateRandomWordSequence").Return(&r)
		deps = &dependencies.Dependencies{
			Config:    config,
			EventEmit: emit,
			Rnd:       rndMock,
		}
	})
	Context("Given a RegisterP2pHandler", func() {
		It("expected 400 for missing OneTimePass", func() {
			router := server.NewRouter(deps)
			w := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: MultiAddr,
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
		It("expected 400 for missing MultiAddr", func() {
			router := server.NewRouter(deps)
			w := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				Otp: OTP,
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
		It("expected 202 Accepted - chat mode", func() {
			router := server.NewRouter(deps)
			w := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: MultiAddr,
				Otp:       OTP,
				Mode:      "chat",
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(w, req)
			ack := model.PeerRegistryAckResponse{}
			b, err := ioutil.ReadAll(w.Body)
			json.Unmarshal(b, &ack)
			Expect(err).To(BeNil())
			Expect(ack.AckId).To(Equal(AckId))
			Expect(w.Code).To(Equal(http.StatusAccepted))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
		It("expected 202 Accepted - file mode", func() {
			router := server.NewRouter(deps)
			w := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: MultiAddr,
				Otp:       OTP,
				Mode:      "file",
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(w, req)
			ack := model.PeerRegistryAckResponse{}
			b, err := ioutil.ReadAll(w.Body)
			json.Unmarshal(b, &ack)
			Expect(err).To(BeNil())
			// fmt.Println("@@@", string(b))

			Expect(ack.AckId).To(Equal(AckId))
			Expect(w.Code).To(Equal(http.StatusAccepted))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
