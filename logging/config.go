package logging

import (
	"github.com/ferdoran/go-sro-framework/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	UnhandledPacketsLogfile = "unhandled_packets.txt"
)

func Init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-01 15:04:05.000",
	})
	logLevel, err := log.ParseLevel(viper.GetString(config.LogLevel))

	if err != nil {
		log.Warnf("failed to parse log level: [%s]. setting to info", logLevel)
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)

	env := viper.GetString(config.Environment)
	if env == "dev" {
		log.Infoln("detected ENV=dev. Setting output to stdout")
		log.SetOutput(os.Stdout)
	} else {
		logFile, err := os.Create(viper.GetString(config.LogFile))
		if err != nil {
			log.Error(err)
		} else {
			log.SetOutput(logFile)
		}
	}
}

func UnhandledPacketLogger() *log.Logger {
	unhandledPacketsLog, err := os.Create(UnhandledPacketsLogfile)
	if err != nil {
		log.Error(err)
	}
	logger := log.New()
	logger.SetOutput(unhandledPacketsLog)
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-01 15:04:05.000",
	})
	return logger
}
