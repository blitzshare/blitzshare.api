package routes

import (
	"blitzshare.api/app/model"
	"net/http"

	"blitzshare.api/app/dependencies"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// DeleteP2pPeerRegistry godoc
// @Summary  Returns peer registry config via OTP
// @Param        X-Api-Key  header  string  true  "API key authenication header"
// @Schemes
// @Tags    p2p-registry
// @Param    otp    path     string  true  "otp obtained whern registrated"
// @Param    token    path     string  true  "deregistration token obtained whern registrated"
// @Success 201 {object} model.AckResponse "deregistration request accepted"
// @Success 404 "otp not found"
// @Success 500 "internal server error"
// @Router   /p2p/registry/{otp}/{token} [delete]
func DeleteP2pPeerHandler(deps *dependencies.Dependencies) func(c *gin.Context) {
	return func(c *gin.Context) {
		AddDefaultResponseHeaders(c)
		if IsNotAuthorized(c, deps.ApiKeychain) {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		token := c.Params.ByName("token")
		otp := c.Params.ByName("otp")
		log.Debugln("GetP2pPeerHandler", token)
		if len(otp) < 10 || len(token) < 10 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		e := &model.P2pPeerDeregisterCmd{
			Otp: model.Otp{
				Otp: otp,
			},
			Token: model.Token{
				Token: token,
			},
		}

		msgId, err := deps.EventEmit.EmitP2pPeerDeregisterCmd(deps.Config.Settings.QueueUrl, deps.Config.ClientId, e)
		if msgId == nil || err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusAccepted, model.AckResponse{AckId: *msgId})
		}
	}
}
