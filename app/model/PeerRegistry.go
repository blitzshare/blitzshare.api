package model

type P2pPeerRegistryReq struct {
	MultiAddr string `binding:"required" json:"multiAddr"`
	Otp       string `binding:"required" json:"otp"`
	Mode      string `binding:"required" json:"mode"`
}

type P2pPeerRegistryResponse struct {
	MultiAddr string `binding:"required" json:"multiAddr"`
	Otp       string `binding:"required" json:"otp"`
	Mode      string `binding:"required" json:"mode"`
}

type P2pPeerRegistryCmd struct {
	MultiAddr string `binding:"required" json:"multiAddr"`
	Otp       string `binding:"required" json:"otp"`
	Mode      string `binding:"required" json:"mode"`
	Token     string `binding:"required" json:"token"`
}

type P2pPeerDeregisterCmd struct {
	Otp   string `binding:"required" json:"otp"`
	Token string `binding:"required" json:"token"`
}

type MultiAddrResponse struct {
	MultiAddr string `binding:"required" json:"multiAddr"`
}

type PeerRegistryAckResponse struct {
	AckId string `binding:"required" json:"ackId"`
	Token string `binding:"required" json:"token"`
}

type AckResponse struct {
	AckId string `binding:"required" json:"ackId"`
}

type NodeConfig struct {
	NodeId string `json:"nodeId"`
	Port   int    `json:"port"`
}

type NodeConfigRespone struct {
	NodeId string `json:"nodeId"`
	Port   int    `json:"port"`
}
