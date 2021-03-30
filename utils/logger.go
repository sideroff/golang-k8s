package utils

import ( 
	log "github.com/sirupsen/logrus"
)

func ConfigureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
}