package test

import (
	"blitzshare.api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"

	. "github.com/onsi/gomega"
)

func MatchAny(input interface{}) bool {
	return true
}

func AsserBlitzshareHeaders(rec *httptest.ResponseRecorder) {
	Expect(rec.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
}

func MockApiKeychain(isValid bool) *mocks.ApiKeychain {
	keychain := &mocks.ApiKeychain{}
	keychain.On("IsValid", mock.MatchedBy(MatchAny)).Return(isValid)
	return keychain
}
