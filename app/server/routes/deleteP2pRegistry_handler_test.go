package routes_test

import (
	"blitzshare.api/app/server/routes"
	"errors"
	"fmt"
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

var _ = Describe("DELETE P2p Registry", func() {
	const (
		OTP             = "gelandelaufer-astromancer-scurvyweed-sayability"
		DeregisterToken = "sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"
	)
	var deps *dependencies.Dependencies
	BeforeEach(func() {
		AckId := "fkjldsfkksFSDFODSKFSJFsdfksdfldskSDFK"
		settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
		emit := &mocks.EventEmit{}
		emit.On("EmitP2pPeerDeregisterCmd",
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
			mock.MatchedBy(test.MatchAny),
		).Return(&AckId, nil)
		config := config.Config{
			Server:   config.Server{Port: 323},
			Settings: settings,
		}
		deps = &dependencies.Dependencies{
			Config:    config,
			EventEmit: emit,
		}
	})

	Context("Given a RegisterP2pHandler", func() {
		It("expected 202 for valid otp and deregister token", func() {
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/p2p/registry/%s/%s", OTP, DeregisterToken), nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusAccepted))
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
				Config:    deps.Config,
				EventEmit: emit,
			}
			router := routes.NewRouter(deps)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/p2p/registry/%s/%s", OTP, DeregisterToken), nil)
			router.ServeHTTP(rec, req)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
