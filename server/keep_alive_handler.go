package server

type KeepAliveHandler struct{}

func NewKeepAliveHandler() PacketHandler {
	handler := KeepAliveHandler{}
	PacketManagerInstance.RegisterHandler(0x2002, handler)
	return handler
}

func (h KeepAliveHandler) Handle(packet PacketChannelData) {
	// Do Nothing
}
