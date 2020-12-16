package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	LogFile                 = "log.txt"
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
	env := os.Getenv("SRO_ENV")
	if env == "dev" {
		log.Infoln("Detected SRO_ENV=dev. Setting log level to debug")
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	} else {
		logFile, err := os.Create(LogFile)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(log.InfoLevel)
		log.SetOutput(logFile)
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
