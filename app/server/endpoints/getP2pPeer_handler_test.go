package endpoints_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/endpoints"
	"blitzshare.api/mocks"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Get P2p Peer address by OTP", func() {
	const (
		OTP       = "gelandelaufer-astromancer-scurvyweed-sayability"
		MultiAddr = "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"
	)
	Context("Given a GetP2pPeerHandler", func() {
		var config = config.Config{
			Server:   config.Server{Port: 323},
			Settings: config.Settings{RedisUrl: "redis.svc.cluster.local"},
		}
		It("expected 400 NotFound", func() {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			registry := &mocks.Registry{}
			registry.On("GetOtp",
				mock.MatchedBy(func(input interface{}) bool {
					return input == "non-defined-id"
				})).Return("", errors.New("test"))
			c.Params = []gin.Param{
				{
					Key:   "otp",
					Value: "non-defined-id",
				},
			}
			rndMock := &mocks.Rnd{}
			r := "random-string"
			rndMock.On("GenerateRandomWordSequence").Return(&r)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
				Rnd:      rndMock,
			}
			endpoints.GetP2pPeerHandler(&deps)(c)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})

		It("expected 200 Ok for valid OTP", func() {
			registry := &mocks.Registry{}
			resp := model.P2pPeerRegistryResponse{
				MultiAddr: MultiAddr,
				Otp:       OTP,
				Mode:      "chat",
			}
			bStr, err := json.Marshal(resp)
			registry.On("GetOtp",
				mock.MatchedBy(func(input interface{}) bool {
					return input == OTP
				})).Return(string(bStr), nil)
			rndMock := &mocks.Rnd{}
			r := "random-string"
			rndMock.On("GenerateRandomWordSequence").Return(&r)
			deps := dependencies.Dependencies{
				Config:   config,
				Registry: registry,
				Rnd:      rndMock,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{
				{
					Key:   "otp",
					Value: OTP,
				},
			}
			endpoints.GetP2pPeerHandler(&deps)(c)

			body, _ := ioutil.ReadAll(w.Body)
			peerInfo := model.MultiAddrResponse{}
			err := json.Unmarshal(body, &peerInfo)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(peerInfo.MultiAddr).To(Equal(MultiAddr))
			Expect(w.Header().Get("X-Blitzshare-Service")).To(Equal("blitzshare.api"))
		})
	})
})
