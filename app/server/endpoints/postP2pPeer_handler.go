package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"

	"blitzshare.api/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	log "github.com/sirupsen/logrus"
)

func RegisterP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		var req model.P2pPeerRegistryReq
		token := *deps.Rnd.GenerateRandomWordSequence()
		if err := c.ShouldBindWith(&req, binding.JSON); err == nil {
			log.Debugln("RegisterP2pPeer", req.Otp)
			e := model.P2pPeerRegistryCmd{
				MultiAddr: req.MultiAddr,
				Mode:      req.Mode,
				Otp:       req.Otp,
				Token:     token,
			}
			msgId, err := deps.EventEmit.EmitP2pPeerRegistryCmd(deps.Config.Settings.QueueUrl, deps.Config.ClientId, &e)
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
			} else {
				rep := model.PeerRegistryAckResponse{
					AckId: msgId,
					Token: token,
				}
				c.JSON(http.StatusAccepted, rep)
			}
		} else {
			log.Errorln("RegisterP2pPeer", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
