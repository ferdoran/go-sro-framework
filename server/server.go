package server

import (
	"fmt"
	"github.com/ferdoran/go-sro-framework/network"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	host              string
	port              int
	Sessions          map[string]*Session
	options           network.EncodingOptions
	ModuleID          string
	PacketChannel     chan PacketChannelData
	BackendConnection chan BackendConnectionData
	SessionClosed     chan *Session
	SessionCreated    chan string
}

func NewEngine(host string, port int, options network.EncodingOptions) Server {
	packetChannel := make(chan PacketChannelData)

	return Server{
		host,
		port,
		make(map[string]*Session),
		options,
		"SampleServer",
		packetChannel,
		make(chan BackendConnectionData, 1),
		make(chan *Session, 8),
		make(chan string, 8),
	}
}

func (e *Server) Start() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		log.Infof("%v signal received. Closing all connections ...", sig.String())
		for _, conn := range e.Sessions {
			if conn != nil && conn.Conn != nil {
				conn.Conn.Close()
			}
		}
		os.Exit(1)
	}()

	NewKeyExchangeHandler()
	NewKeyExchangeCompletedHandler()
	NewModuleIdentifactionHandler()
	NewKeepAliveHandler()

	log.Infof("Started listening on %s:%d\n", e.host, e.port)
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", e.host, e.port))
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()

		log.Infof("New connection from %v\n", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}

		session := NewSession(
			conn,
			e.options,
			e.PacketChannel,
			e.SessionClosed,
			e.BackendConnection,
			e.ModuleID,
		)
		e.Sessions[session.ID] = session
		go session.StartHandling()
		e.SessionCreated <- session.ID
	}
}
