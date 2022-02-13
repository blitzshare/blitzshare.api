package model

type MultiAddr struct {
	MultiAddr string `binding:"required" json:"multiAddr" example:"/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"`
}
type Otp struct {
	Otp string `binding:"required" json:"otp"  example:"gelandelaufer-astromancer-scurvyweed-sayability"`
}

type Mode struct {
	Mode string `binding:"required" json:"mode" example:"file"`
}

type Token struct {
	Token string `binding:"required" json:"token" example:"sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"`
}
type P2pPeerRegistryReq struct {
	MultiAddr
	Otp
	Mode string `binding:"required" json:"mode" example:"chat"`
}

type P2pPeerRegistryResponse struct {
	MultiAddr
	Otp
	Mode
}

type P2pPeerRegistryCmd struct {
	MultiAddr
	Otp
	Mode
	Token
}

type P2pPeerDeregisterCmd struct {
	Otp
	Token
}

type MultiAddrResponse struct {
	MultiAddr
}

type PeerRegistryAckResponse struct {
	AckResponse
	Token
}

type AckResponse struct {
	AckId string `binding:"required" json:"ackId" example:"12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"`
}

type NodeConfig struct {
	NodeId string `json:"nodeId" example:"234DSFG345dsfgdfghdfgdfgDFGASR4"`
	Port   int    `json:"port" example:"6789"`
}

type NodeConfigRespone struct {
	NodeConfig
}
