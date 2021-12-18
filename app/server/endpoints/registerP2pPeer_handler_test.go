package endpoints_test

import (
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Register P2p Endpoint", func() {
	var deps *dependencies.Dependencies
	BeforeEach(func() {
		server := config.Server{Port: 323}
		settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
		config := config.Config{
			Server:   server,
			Settings: settings,
		}
		deps, _ = dependencies.NewDependencies(config)
	})
	Context("Given a RegisterP2pHandler", func() {
		It("expected 400", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			endpoints.RegisterP2pPeerHandler(deps)(c)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
