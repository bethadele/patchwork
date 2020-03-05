package main

import (
	"net/http"

	"github.com/bethadele/patchwork/server"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "0.1"

	listenAddr = kingpin.Flag("listenaddr", "Address for the server to listen on.").Short('l').Default("0.0.0.0:8080").String()
	logDebug   = kingpin.Flag("debug", "Enable debug level logging.").Short('d').Bool()
)

func init() {
	kingpin.Version(version)
	kingpin.Parse()

	if *logDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	server := server.New()
	if err := http.ListenAndServe(*listenAddr, server); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
