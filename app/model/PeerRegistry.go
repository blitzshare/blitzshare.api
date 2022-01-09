package model

type P2pPeerRegistryCmd struct {
	MultiAddr string `form:"multiAddr" binding:"required" json:"multiAddr"`
	Otp       string `form:"otp" binding:"required" json:"otp"`
}

type MultiAddrResponse struct {
	MultiAddr string `binding:"required" json:"multiAddr"`
}

type PeerRegistryAckResponse struct {
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
