package routes_test

import (
	"blitzshare.api/app/model"
	"blitzshare.api/app/routes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/mocks"
	"blitzshare.api/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("POST /p2p/registry", func() {
	const (
		OTP       = "gelandelaufer-astromancer-scurvyweed-sayability"
		MultiAddr = "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"
		AckId     = "12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"
	)
	var deps *dependencies.Dependencies
	BeforeEach(func() {
		server := config.Server{Port: 323}
		AckId := "12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"
		settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
		emit := &mocks.EventEmit{}
		emit.On("EmitP2pPeerRegisterCmd",
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
		).Return(&AckId, nil)
		config := config.Config{
			Server:   server,
			Settings: settings,
		}
		rndMock := &mocks.Rnd{}
		r := "random-string"
		rndMock.On("GenerateRandomWordSequence").Return(&r)
		deps = &dependencies.Dependencies{
			Config:      config,
			EventEmit:   emit,
			Rnd:         rndMock,
			ApiKeychain: test.MockApiKeychain(true),
		}
	})
	Context("PostP2pRegistryHandler", func() {
		It("expected 400 for missing OneTimePass", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: model.MultiAddr{
					MultiAddr: MultiAddr,
				},
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 400 for missing MultiAddr", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				Otp: model.Otp{
					Otp: OTP,
				},
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 202 Accepted - chat mode", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: model.MultiAddr{
					MultiAddr: MultiAddr,
				},
				Otp: model.Otp{
					Otp: OTP,
				},
				Mode: "chat",
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(rec, req)
			ack := model.PeerRegistryAckResponse{}
			b, err := ioutil.ReadAll(rec.Body)
			json.Unmarshal(b, &ack)
			Expect(err).To(BeNil())
			Expect(ack.AckId).To(Equal(AckId))
			Expect(rec.Code).To(Equal(http.StatusAccepted))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 202 Accepted - file mode", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			body, _ := json.Marshal(model.P2pPeerRegistryReq{
				MultiAddr: model.MultiAddr{
					MultiAddr: MultiAddr,
				},
				Otp: model.Otp{
					Otp: OTP,
				},
				Mode: "file",
			})
			req, _ := http.NewRequest("POST", "/p2p/registry", bytes.NewReader(body))
			router.ServeHTTP(rec, req)
			ack := model.PeerRegistryAckResponse{}
			b, err := ioutil.ReadAll(rec.Body)
			json.Unmarshal(b, &ack)
			Expect(err).To(BeNil())
			Expect(ack.AckId).To(Equal(AckId))
			Expect(rec.Code).To(Equal(http.StatusAccepted))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
