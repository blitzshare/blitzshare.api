package model

type PeerRegistry struct {
	MultiAddr   string `form:"multiAddr" binding:"required"`
	OneTimePass string `form:"oneTimePass" binding:"required"`
}

