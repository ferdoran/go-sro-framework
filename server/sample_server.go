package server

import (
	"github.com/ferdoran/go-sro-framework/config"
	"github.com/ferdoran/go-sro-framework/network"
	log "github.com/sirupsen/logrus"
)

func StartSampleServer() {

	server := NewEngine(
		[]byte{127, 0, 0, 1},
		15779,
		network.EncodingOptions{
			None:         false,
			Disabled:     false,
			Encryption:   true,
			EDC:          true,
			KeyExchange:  true,
			KeyChallenge: false,
		},
		config.Config{
			GatewayServer: config.GatewayServerGlobalConfig{},
			DB: struct {
				Account config.DBConfig `json:"account"`
				Shard   config.DBConfig `json:"shard"`
			}{},
			AgentServer: config.AgentServerGlobalConfig{},
		},
	)
	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
