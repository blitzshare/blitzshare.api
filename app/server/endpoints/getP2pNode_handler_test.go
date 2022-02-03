package endpoints_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"

	"net/http"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/endpoints"
	"blitzshare.api/mocks"
	"blitzshare.api/test"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get /p2p/bootstrap-node", func() {
	Context("Given GetP2pBootstrapNodeHandler", func() {
		var config = config.Config{
			Server:   config.Server{Port: 323},
			Settings: config.Settings{RedisUrl: "redis.svc.cluster.local"},
		}
		It("expected 400 NotFound", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			endpoints.GetP2pBootstrapNodeHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 500 InternalServerError", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("not-parsable-id", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			endpoints.GetP2pBootstrapNodeHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			test.AsserBlitzshareHeaders(rec)
		})

		It("expected 200 Ok for valid OTP", func() {
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("{\"nodeId\":\"note-test-id\",\"port\":63785}", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			endpoints.GetP2pBootstrapNodeHandler(&deps)(c)

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
