package routes

import (
	"blitzshare.api/app/model"
	"net/http"

	"blitzshare.api/app/dependencies"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	log "github.com/sirupsen/logrus"
)

// PostP2pRegistry godoc
// @Summary  Registers peer registry config
// @Param        X-Api-Key  header  string  true  "API key authenication header"
// @Schemes
// @Tags    p2p-registry
// @Param   config body     model.P2pPeerRegistryReq  true  "p2p registry config"
// @Success 202 {object} model.PeerRegistryAckResponse "acknowledge response with de-registration token"
// @Success 500 "failed to register peer"
// @Success 400 "invalid params"
// @Router   /p2p/registry [post]
func PostP2pRegistryHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		if IsNotAuthorized(c, deps.ApiKeychain) {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		AddDefaultResponseHeaders(c)
		var req model.P2pPeerRegistryReq
		token := *deps.Rnd.GenerateRandomWordSequence()
		if err := c.ShouldBindWith(&req, binding.JSON); err == nil {
			log.Debugln("RegisterP2pPeer", req.Otp)
			e := model.P2pPeerRegistryCmd{
				MultiAddr: req.MultiAddr,
				Mode: model.Mode{
					Mode: req.Mode,
				},
				Otp: req.Otp,
				Token: model.Token{
					Token: token,
				},
			}
			msgId, err := deps.EventEmit.EmitP2pPeerRegisterCmd(deps.Config.Settings.QueueUrl, deps.Config.ClientId, &e)
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
			} else {
				rep := model.PeerRegistryAckResponse{
					AckResponse: model.AckResponse{
						AckId: *msgId,
					},
					Token: model.Token{
						Token: token,
					},
				}
				c.JSON(http.StatusAccepted, rep)
			}
		} else {
			log.Errorln("RegisterP2pPeer", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
