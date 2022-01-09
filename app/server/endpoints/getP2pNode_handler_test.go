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
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get P2p Bootstrap Node", func() {
	Context("Given a P2p Bootstrap registered in the system", func() {
		var config = config.Config{
			Server:   config.Server{Port: 323},
			Settings: config.Settings{RedisUrl: "redis.svc.cluster.local"},
		}
		It("expected 400 NotFound", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			endpoints.GetP2pBootstrapNode(&deps)(c)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
		It("expected 500 InternalServerError", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("not-parsable-id", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			endpoints.GetP2pBootstrapNode(&deps)(c)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})

		It("expected 200 Ok for valid OTP", func() {
			registry := &mocks.Registry{}
			registry.On("GetNodeConfig").Return("{\"nodeId\":\"note-test-id\",\"port\":63785}", nil)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			endpoints.GetP2pBootstrapNode(&deps)(c)

			body, _ := ioutil.ReadAll(w.Body)
			peerInfo := model.NodeConfigRespone{}
			err := json.Unmarshal(body, &peerInfo)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(peerInfo.NodeId).To(Equal("note-test-id"))
			Expect(peerInfo.Port).To(Equal(63785))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
