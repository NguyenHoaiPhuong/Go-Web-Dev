package log

import (
	"io/ioutil"

	stdHook "github.com/NguyenHoaiPhuong/Go-Web-Dev/012_LoadBalancer/log/hooks/std"

	"github.com/sirupsen/logrus"
)

// ILog interface includes methods related to logger
type ILog interface {
	addSTDHook(int)

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})
}

// Log struct
type Log struct {
	ILog

	tmHook *stdHook.Hook
	logger *logrus.Logger
}

// init : initializes the logger
func (l *Log) init() {
	// Init logger as below
	l.logger = logrus.New()

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	l.logger.SetOutput(ioutil.Discard)

	// Log as JSON instead of the default ASCII formatter.
	// l.logger.SetFormatter(&log.JSONFormatter{})
	l.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	l.logger.SetLevel(logrus.TraceLevel)
}

// addSTDHook : initialize terminal hook
func (l *Log) addSTDHook(level int) {
	// Allow only 1 standard hook
	if l.tmHook == nil {
		l.tmHook = new(stdHook.Hook)
		l.tmHook.Init(level)
		l.logger.AddHook(l.tmHook)
	}
}

// Panic : panic msg
func (l *Log) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Panicf : panic msg with format
func (l *Log) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// Panicln : panic msg + new line
func (l *Log) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Fatal : fatal msg
func (l *Log) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf : fatal msg with format
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Fatalln : fatal msg + new line
func (l *Log) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

// Error : error msg
func (l *Log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorf : error msg with format
func (l *Log) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Errorln : error msg + new line
func (l *Log) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// Warn : warn msg
func (l *Log) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnf : warn msg with format
func (l *Log) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Warnln : warn msg + new line
func (l *Log) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Info : info msg
func (l *Log) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infof : info msg with format
func (l *Log) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Infoln : info msg + new line
func (l *Log) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Debug : debug msg
func (l *Log) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugf : debug msg with format
func (l *Log) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Debugln : debug msg + new line
func (l *Log) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Trace : trace msg
func (l *Log) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

// Tracef : trace msg with format
func (l *Log) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

// Traceln : trace msg + new line
func (l *Log) Traceln(args ...interface{}) {
	l.logger.Traceln(args...)
}
