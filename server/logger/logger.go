package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// CtxLogger holds config for the rest of the application
var CtxLogger = log.New()

func init() {
	CtxLogger.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	CtxLogger.SetOutput(os.Stdout)

	// Can be used with environment variable or flag in future
	CtxLogger.SetLevel(log.DebugLevel)
}
