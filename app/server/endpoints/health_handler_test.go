package endpoints_test

import (
	"net/http/httptest"
	"testing"

	"github.com/blitzshare/blitzshare.fileshare.api/app/server/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	endpoints.HealthHandler()(c)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "im alive", w.Body.String())
}
