package server

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type PacketManager struct {
	Handlers map[uint16]PacketHandler
	mutex    sync.Mutex
}

func (pm *PacketManager) RegisterHandler(opcode uint16, handler PacketHandler) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	if pm.Handlers[opcode] != nil {
		logrus.Debugf("packet handler for opcode %X already registered\n", opcode)
		return
	}
	pm.Handlers[opcode] = handler
}

var PacketManagerInstance = &PacketManager{Handlers: make(map[uint16]PacketHandler)}
