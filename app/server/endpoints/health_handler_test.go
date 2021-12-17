package endpoints_test

import (
	"net/http/httptest"

	"blitzshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Register P2p Endpoint", func() {
	Context("Given a RegisterP2pHandler", func() {
		It("expected 200 OK", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			endpoints.HealthHandler()(c)
			Expect(w.Code).To(Equal(200))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
