package simplelog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type SimpleLogger struct {
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

func Logger(level string) *SimpleLogger {

	var logSetup *SimpleLogger

	if strings.EqualFold(level, "DEBUG") {
		logSetup = setup(os.Stdout, os.Stdout, os.Stderr)
	} else if strings.EqualFold(level, "ERROR") {
		logSetup = setup(ioutil.Discard, ioutil.Discard, os.Stderr)
	} else if strings.EqualFold(level, "NONE") {
		logSetup = setup(ioutil.Discard, ioutil.Discard, ioutil.Discard)
	} else {
		logSetup = setup(ioutil.Discard, os.Stdout, os.Stderr)
		if !strings.EqualFold(level, "INFO") {
			logSetup.Error.Printf("Logger: level \"%s\" unrecognised. Defaulting to \"%s\"", level, "INFO")
		}
	}
	return logSetup
}

func setup(debugHandle io.Writer, infoHandle io.Writer, errorHandle io.Writer) *SimpleLogger {

	logger := SimpleLogger{}

	logger.Debug = log.New(debugHandle,
		"DEB: ",
		log.Ldate|log.Lmicroseconds|log.LUTC)

	logger.Info = log.New(infoHandle,
		"INF: ",
		log.Ldate|log.Lmicroseconds|log.LUTC)

	logger.Error = log.New(errorHandle,
		"ERR: ",
		log.Ldate|log.Lmicroseconds|log.LUTC)

	logger.Debug.Printf("Debug logging enabled")

	return &logger

}
