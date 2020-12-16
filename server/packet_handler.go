package server

type PacketHandler interface {
	Handle(packet PacketChannelData)
}
