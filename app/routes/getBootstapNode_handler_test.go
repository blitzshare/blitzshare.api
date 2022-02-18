package routes_test

import (
	"blitzshare.api/app/model"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"

	"net/http"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/routes"
	"blitzshare.api/mocks"
	"blitzshare.api/test"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get /p2p/bootstrap-node", func() {
	Context("Given GetBootstrapNodeHandler", func() {
		var config = config.Config{
			Server:   config.Server{Port: 323},
			Settings: config.Settings{RedisUrl: "redis.svc.cluster.local"},
		}
		It("expected 401 Unauthorized", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("", nil)
			keychain := &mocks.ApiKeychain{}
			keychain.On("IsValid", mock.MatchedBy(test.MatchAny)).Return(false)
			deps := dependencies.Dependencies{
				Config:      config,
				Registry:    registry,
				ApiKeychain: keychain,
			}
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Request.Header.Set("x-api-key", "test")
			routes.GetBootstrapNodeHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusUnauthorized))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 400 NotFound", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("", nil)
			keychain := &mocks.ApiKeychain{}
			keychain.On("IsValid", mock.MatchedBy(test.MatchAny)).Return(true)
			deps := dependencies.Dependencies{
				Config:      config,
				Registry:    registry,
				ApiKeychain: keychain,
			}
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Request.Header.Set("x-api-key", "test")
			routes.GetBootstrapNodeHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 500 InternalServerError", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			registry := &mocks.Registry{}
			keychain := &mocks.ApiKeychain{}
			keychain.On("IsValid", mock.MatchedBy(test.MatchAny)).Return(true)
			registry.On("GetNodeConfig").Return("not-parsable-id", nil)
			deps := dependencies.Dependencies{
				Config:      config,
				Registry:    registry,
				ApiKeychain: keychain,
			}
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Request.Header.Set("x-api-key", "test")
			routes.GetBootstrapNodeHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 200 Ok for valid OTP", func() {
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("{\"nodeId\":\"note-test-id\",\"port\":63785}", nil)
			keychain := &mocks.ApiKeychain{}
			keychain.On("IsValid", mock.MatchedBy(test.MatchAny)).Return(true)
			deps := dependencies.Dependencies{
				Config:      config,
				Registry:    registry,
				ApiKeychain: keychain,
			}
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Request.Header.Set("x-api-key", "test")
			routes.GetBootstrapNodeHandler(&deps)(c)

			body, _ := ioutil.ReadAll(rec.Body)
			peerInfo := model.NodeConfigRespone{}
			err := json.Unmarshal(body, &peerInfo)
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(peerInfo.NodeId).To(Equal("note-test-id"))
			Expect(peerInfo.Port).To(Equal(63785))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
