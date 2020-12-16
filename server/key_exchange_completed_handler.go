package server

import (
	log "github.com/sirupsen/logrus"
)

type KeyExchangeCompletedHandler struct {
}

func NewKeyExchangeCompletedHandler() PacketHandler {
	handler := KeyExchangeCompletedHandler{}
	PacketManagerInstance.RegisterHandler(0x9000, handler)
	return handler
}

func (h KeyExchangeCompletedHandler) Handle(packet PacketChannelData) {
	if !packet.Session.Context.StartedHandshake && !packet.Session.Context.CompletedLocalSetup && !packet.Session.Context.CompletedRemoteSetup {
		log.Error("Handshake still not setup completely")
	}
	log.Debugln("Finished Handshake")
	packet.Session.Context.FinishedHandshake = true
}
