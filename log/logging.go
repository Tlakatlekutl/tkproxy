package log

import (
	"strings"
	"log"
	"io"
)

type DebugLevel int

const (
	err DebugLevel = 1 << iota
	warning
	info
	trace
)

var DLevel DebugLevel = err

// Calculate debug level based on input arguments.
// default ERROR
//
// Example:
// SetDebugLevel("ERROR", "TRACE")
// SetDebugLevel("ALL")
// SetDebugLevel("NOLOG")
func SetDebugLevel(args ...string) bool {
	for _, l := range args {
		switch strings.ToUpper(l) {
		case "ERROR":
			DLevel |= err
		case "WARNING":
			DLevel |= warning
		case "INFO":
			DLevel |= info
		case "TRACE":
			DLevel |= trace
		case "NOLOG":
			DLevel = 0
		case "ALL":
			DLevel = err | warning | info | trace
		default:
			Error("Unknown debug level: %s", l)
			return false
		}
	}
	return true
}

// Set io.Writer for output
// default is "stdout"
func SetLogOUT(w io.Writer) {
	log.SetOutput(w)
}

// Log a message if the ERROR parameter is enabled
func Error(msg string, args ...interface{})  {
	if DLevel & err == err {
		msg = "[ERROR] "+ msg
		log.Printf(msg, args...)
	}
}

// Log a message if the WARNING parameter is enabled
func Warning(msg string, args ...interface{})  {
	if DLevel & warning == warning {
		msg = "[WARN ] "+ msg
		log.Printf(msg, args...)
	}
}

// Log a message if the INFO parameter is enabled
func Info(msg string, args ...interface{})  {
	if DLevel & info == info {
		msg = "[INFO ] "+ msg
		log.Printf(msg, args...)
	}
}

// Log a message if the TRACE parameter is enabled
func Trace(msg string, args ...interface{})  {
	if DLevel & trace == trace {
		msg = "[TRACE] "+ msg
		log.Printf(msg, args...)
	}
}
