package endpoints

import (
	"net/http"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DeleteP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		token := c.Params.ByName("token")
		otp := c.Params.ByName("otp")
		log.Debugln("GetP2pPeerHandler", token)
		if len(otp) < 10 || len(token) < 10 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		e := &model.P2pPeerDeregisterCmd{
			Otp:   otp,
			Token: token,
		}

		msgId, err := deps.EventEmit.EmitP2pPeerDeregisterCmd(deps.Config.Settings.QueueUrl, deps.Config.ClientId, e)
		if msgId != nil || err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusAccepted, gin.H{"ackId": msgId})
		}
	}
}
