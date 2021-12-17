package endpoints_test

import (
	"net/http/httptest"

	"net/http"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get P2p Peer address by OTP", func() {
	Context("Given a GetP2pPeerHandler", func() {
		It("expected 400", func() {
			server := config.Server{Port: 323}
			settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
			config := config.Config{
				Server:   server,
				Settings: settings,
			}
			deps, _ := dependencies.NewDependencies(config)
			// deps := dependencies.Dependencies{
			// 	Config: config,
			// 	Registry: ?,
			// }
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: "non-defined-id",
				},
			}
			endpoints.GetP2pPeerHandler(deps)(c)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
