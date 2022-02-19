package routes_test

import (
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/routes"
	"blitzshare.api/test"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health check test", func() {
	Context("Given a healthcheck", func() {
		It("expected 200 OK", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			routes.HealthHandler()(c)
			Expect(rec.Code).To(Equal(http.StatusOK))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
