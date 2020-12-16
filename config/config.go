package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strings"
)

type Config struct {
	GatewayServer GatewayServerGlobalConfig `json:"gw"`
	DB            struct {
		Account DBConfig `json:"account"`
		Shard   DBConfig `json:"shard"`
	} `json:"db"`
	AgentServer AgentServerGlobalConfig `json:"agent"`
}

var GlobalConfig Config

func LoadConfig(configFile string) {
	logrus.Printf("loading config: %s\n", configFile)
	file, err := os.Open(configFile)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)

	if err != nil {
		panic(err.Error())
	}

	GlobalConfig = cfg
}

// "user:password@tcp(127.0.0.1:3306)/hello"
type DBConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Params   string `json:"params"`
}

func ConnStringAccount() string {
	var sb strings.Builder
	sb.WriteString(GlobalConfig.DB.Account.User)
	sb.WriteString(":")
	sb.WriteString(GlobalConfig.DB.Account.Password)
	sb.WriteString("@tcp(")
	sb.WriteString(GlobalConfig.DB.Account.Host)
	sb.WriteString(":")
	sb.WriteString(GlobalConfig.DB.Account.Port)
	sb.WriteString(")/")
	sb.WriteString(GlobalConfig.DB.Account.Database)
	sb.WriteString(GlobalConfig.DB.Account.Params)
	return sb.String()
}

func ConnStringShard() string {
	var sb strings.Builder
	sb.WriteString(GlobalConfig.DB.Shard.User)
	sb.WriteString(":")
	sb.WriteString(GlobalConfig.DB.Shard.Password)
	sb.WriteString("@tcp(")
	sb.WriteString(GlobalConfig.DB.Shard.Host)
	sb.WriteString(":")
	sb.WriteString(GlobalConfig.DB.Shard.Port)
	sb.WriteString(")/")
	sb.WriteString(GlobalConfig.DB.Shard.Database)
	sb.WriteString(GlobalConfig.DB.Shard.Params)
	return sb.String()
}

type GatewayServerGlobalConfig struct {
	Port     int    `json:"port"`
	IP       net.IP `json:"ip"`
	ModuleID string `json:"moduleId"`
	Secret   string `json:"secret"`
}

type AgentServerGlobalConfig struct {
	Port       int    `json:"port"`
	IP         net.IP `json:"ip"`
	ModuleID   string `json:"moduleId"`
	Secret     string `json:"secret"`
	DataPath   string `json:"dataPath"`   // Path to extracted Data.pk2
	NavmeshGOB string `json:"navmeshGob"` // Path to prelinked navmesh data as GOB
}
