package routes_test

import (
	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/routes"
	"blitzshare.api/mocks"
	"blitzshare.api/test"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("DELETE P2p Registry", func() {
	const (
		OTP             = "gelandelaufer-astromancer-scurvyweed-sayability"
		DeregisterToken = "sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"
		AckId           = "fkjldsfkksFSDFODSKFSJFsdfksdfldskSDFK"
	)
	var deps *dependencies.Dependencies
	BeforeEach(func() {
		settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
		emit := &mocks.EventEmit{}
		ack := AckId
		emit.On("EmitP2pPeerDeregisterCmd",
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
		).Return(&ack, nil)
		config := config.Config{
			Server:   config.Server{Port: 323},
			Settings: settings,
		}
		deps = &dependencies.Dependencies{
			Config:      config,
			EventEmit:   emit,
			ApiKeychain: test.MockApiKeychain(true),
		}
	})

	Context("Given a RegisterP2pHandler", func() {
		It("expected 202 for valid otp and deregister token", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/p2p/registry/%s/%s", OTP, DeregisterToken)
			req, _ := http.NewRequest("DELETE", url, nil)
			req.Header.Set("X-Api-Key", "test")
			router.ServeHTTP(rec, req)
			var response model.AckResponse
			body, _ := ioutil.ReadAll(rec.Body)
			json.Unmarshal(body, &response)
			Expect(rec.Code).To(Equal(http.StatusAccepted))
			Expect(response.AckId).To(Equal(AckId))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 404 for undefined deregistration token", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/p2p/registry/%s", OTP), nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 404 for undefined otp token", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/p2p/registry", nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 400 for les then 1o char token and otp length", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/p2p/registry/123456789/123456789", nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 500 when deregister event fails", func() {
			emit := &mocks.EventEmit{}
			emit.On("EmitP2pPeerDeregisterCmd",
				mock.MatchedBy(test.MatchAny),
				mock.MatchedBy(test.MatchAny),
				mock.MatchedBy(test.MatchAny),
			).Return(nil, errors.New("failed to acknowledge event"))
			deps = &dependencies.Dependencies{
				Config:      deps.Config,
				EventEmit:   emit,
				ApiKeychain: test.MockApiKeychain(true),
			}
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/p2p/registry/%s/%s", OTP, DeregisterToken), nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 401 Unauthorized", func() {
			rec := httptest.NewRecorder()
			deps := dependencies.Dependencies{
				ApiKeychain: test.MockApiKeychain(false),
			}
			router := routes.NewRouter(&deps)
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/p2p/registry/%s/%s", OTP, DeregisterToken), nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusUnauthorized))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
