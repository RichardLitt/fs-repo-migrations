package eventlog

import (
	"bufio"
	"io"
	"os"

	"github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/gopkg.in/natefinch/lumberjack.v2"
)

// init sets up sane defaults
func init() {
	Configure(TextFormatter)
	Configure(Output(os.Stderr))
	// has the effect of disabling logging since we log event entries at Info
	// level by convention
	Configure(LevelError)
}

type Option func()

// Configure applies the provided options sequentially from left to right
func Configure(options ...Option) {
	for _, f := range options {
		f()
	}
}

// LdJSONFormatter Option formats the event log as line-delimited JSON
var LdJSONFormatter = func() {
	logrus.SetFormatter(&PoliteJSONFormatter{})
}

// TextFormatter Option formats the event log as human-readable plain-text
var TextFormatter = func() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

type LogRotatorConfig struct {
	Filename   string
	MaxSizeMB  int
	MaxBackups int
	MaxAgeDays int
}

func Output(w io.Writer) Option {
	return func() {
		logrus.SetOutput(w)
		// TODO return previous Output option
	}
}

func OutputRotatingLogFile(config LogRotatorConfig) Option {
	return func() {
		logrus.SetOutput(
			bufio.NewWriter(&lumberjack.Logger{
				Filename:   config.Filename,
				MaxSize:    int(config.MaxSizeMB),
				MaxBackups: int(config.MaxBackups),
				MaxAge:     int(config.MaxAgeDays),
			}))
	}
}

// LevelDebug Option sets the log level to debug
var LevelDebug = func() {
	logrus.SetLevel(logrus.DebugLevel)
}

// LevelDebug Option sets the log level to error
var LevelError = func() {
	logrus.SetLevel(logrus.ErrorLevel)
}

// LevelDebug Option sets the log level to info
var LevelInfo = func() {
	logrus.SetLevel(logrus.InfoLevel)
}
