package server

import (
	"github.com/ferdoran/go-sro-framework/network"
	log "github.com/sirupsen/logrus"
)

type ModuleIdentificationHandler struct {
}

func NewModuleIdentifactionHandler() PacketHandler {
	handler := ModuleIdentificationHandler{}
	PacketManagerInstance.RegisterHandler(0x2001, handler)
	return handler
}

func (mih ModuleIdentificationHandler) Handle(packet PacketChannelData) {
	moduleName, err := packet.ReadString()
	if err != nil {
		log.Error("Could not read module name")
	}
	isLocalModule, err := packet.ReadByte()
	if err != nil {
		log.Error("Could not read if module is local or not")
	}

	log.Debugf("Module identified %v. Local: %v", moduleName, isLocalModule)

	p := network.EmptyPacket()
	p.MessageID = 0x2001
	p.WriteString(packet.Session.ServerModuleID)
	p.WriteByte(0)
	packet.Session.Conn.Write(p.ToBytes())
}
