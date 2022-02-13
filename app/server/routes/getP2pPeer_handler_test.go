package routes_test

import (
	"blitzshare.api/app/server/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server/routes"
	"blitzshare.api/mocks"
	"blitzshare.api/test"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("GET /p2p/registry/:otp", func() {
	const (
		OTP       = "gelandelaufer-astromancer-scurvyweed-sayability"
		MultiAddr = "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"
	)
	Context("Given GetP2pRegistryHandler", func() {
		var config = config.Config{
			Server:   config.Server{Port: 323},
			Settings: config.Settings{RedisUrl: "redis.svc.cluster.local"},
		}
		It("expected 400 NotFound", func() {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
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
			routes.GetP2pRegistryHandler(&deps)(c)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			test.AsserBlitzshareHeaders(rec)
		})
		It("expected 200 Ok for valid OTP", func() {
			registry := &mocks.Registry{}
			resp := model.P2pPeerRegistryResponse{
				MultiAddr: model.MultiAddr{
					MultiAddr: MultiAddr,
				},
				Otp: model.Otp{
					Otp: OTP,
				},
				Mode: model.Mode{
					Mode: "chat",
				},
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

			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Params = []gin.Param{
				{
					Key:   "otp",
					Value: OTP,
				},
			}
			routes.GetP2pRegistryHandler(&deps)(c)

			body, _ := ioutil.ReadAll(rec.Body)
			peerInfo := model.MultiAddrResponse{}

			err = json.Unmarshal(body, &peerInfo)
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(peerInfo.MultiAddr.MultiAddr).To(Equal(MultiAddr))
			test.AsserBlitzshareHeaders(rec)
		})
	})
})
