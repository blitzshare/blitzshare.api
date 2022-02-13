package model

type P2pPeerRegistryReq struct {
	MultiAddr string `binding:"required" json:"multiAddr" example:"/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"`
	Otp       string `binding:"required" json:"otp"  example:"gelandelaufer-astromancer-scurvyweed-sayability"`
	Mode      string `binding:"required" json:"mode" example:"chat"`
}

type P2pPeerRegistryResponse struct {
	MultiAddr string `binding:"required" json:"multiAddr" example:"/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"`
	Otp       string `binding:"required" json:"otp" example:"gelandelaufer-astromancer-scurvyweed-sayability"`
	Mode      string `binding:"required" json:"mode" example:"file"`
}

type P2pPeerRegistryCmd struct {
	MultiAddr string `binding:"required" json:"multiAddr" example:"/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"`
	Otp       string `binding:"required" json:"otp" example:"gelandelaufer-astromancer-scurvyweed-sayability"`
	Mode      string `binding:"required" json:"mode" example:"file"`
	Token     string `binding:"required" json:"token" example:"sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"`
}

type P2pPeerDeregisterCmd struct {
	Otp   string `binding:"required" json:"otp" example:"gelandelaufer-astromancer-scurvyweed-sayability"`
	Token string `binding:"required" json:"token" example:"sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"`
}

type MultiAddrResponse struct {
	MultiAddr string `binding:"required" json:"multiAddr" example:"/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR"`
}

type PeerRegistryAckResponse struct {
	AckId string `binding:"required" json:"ackId" example:"12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"`
	Token string `binding:"required" json:"token" example:"sdfklsfSDFKSDFmxcvlsdfsdfdfSDFkSDFsdf"`
}

type AckResponse struct {
	AckId string `binding:"required" json:"ackId" example:"12D3KooWQcWw5RGtDqCq43M3t1t43k1CBJ8XPdrU5Bc1KtLnTYK"`
}

type NodeConfig struct {
	NodeId string `json:"nodeId" example:"234DSFG345dsfgdfghdfgdfgDFGASR4"`
	Port   int    `json:"port" example:"6789"`
}

type NodeConfigRespone struct {
	NodeId string `json:"nodeId" example:"234DSFG345dsfgdfghdfgdfgDFGASR4"`
	Port   int    `json:"port" example:"6789"`
}
