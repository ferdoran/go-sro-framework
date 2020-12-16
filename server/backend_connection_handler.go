package server

import (
	"github.com/ferdoran/go-sro-framework/config"
	"github.com/ferdoran/go-sro-framework/network/opcode"
	log "github.com/sirupsen/logrus"
)

type BackendConnectionData struct {
	*Session
	ModuleID string
}

type BackendConnectionHandler struct {
	BackendConnected chan BackendConnectionData
	Config           config.Config
}

func NewBackendConnectionHandler(backendConnectedChannel chan BackendConnectionData, config config.Config) PacketHandler {
	handler := &BackendConnectionHandler{
		BackendConnected: backendConnectedChannel,
		Config:           config,
	}
	PacketManagerInstance.RegisterHandler(opcode.BackendAuthentication, handler)
	return handler
}

func (h *BackendConnectionHandler) Handle(packet PacketChannelData) {
	serverModuleId, err := packet.ReadString()
	if err != nil {
		log.Error("Could not read server name")
	}

	secret, err := packet.ReadString()
	if err != nil {
		log.Error("Could not read secret")
	}

	switch serverModuleId {
	case h.Config.AgentServer.ModuleID:
		if secret != h.Config.AgentServer.Secret {
			packet.Session.Conn.Close()
			log.Error("invalid agent server secret")
		}
	case h.Config.GatewayServer.ModuleID:
		if secret != h.Config.GatewayServer.Secret {
			packet.Session.Conn.Close()
			log.Error("invalid gateway server secret")
		}
	}
	log.Infof("%s connected\n", serverModuleId)

	h.BackendConnected <- BackendConnectionData{
		Session:  packet.Session,
		ModuleID: serverModuleId,
	}
}
