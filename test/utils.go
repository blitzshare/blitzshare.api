package test

import (
	"net/http/httptest"

	. "github.com/onsi/gomega"
)

func MatchAny(input interface{}) bool {
	return true
}

func AsserBlitzshareHeaders(rec *httptest.ResponseRecorder) {
	Expect(rec.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
}
