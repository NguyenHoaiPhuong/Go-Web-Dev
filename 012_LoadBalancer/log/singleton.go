package log

var globalLogger *Log = nil

func init() {
	globalLogger = new(Log)
	globalLogger.init()
}

// Logger returns globalLogger
func Logger() ILog {
	return globalLogger
}

// SetSTDHook : set level for standard hook of the global logger
func SetSTDHook(level int) {
	globalLogger.addSTDHook(level)
}

// Panic : panic msg
func Panic(args ...interface{}) {
	globalLogger.Panic(args...)
}

// Panicf : panic msg with format
func Panicf(format string, args ...interface{}) {
	globalLogger.Panicf(format, args...)
}

// Panicln : panic msg + new line
func Panicln(args ...interface{}) {
	globalLogger.Panicln(args...)
}

// Fatal : fatal msg
func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

// Fatalf : fatal msg with format
func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

// Fatalln : fatal msg + new line
func Fatalln(args ...interface{}) {
	globalLogger.Fatalln(args...)
}

// Error : error msg
func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

// Errorf : error msg with format
func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

// Errorln : error msg + new line
func Errorln(args ...interface{}) {
	globalLogger.Errorln(args...)
}

// Warn : warn msg
func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}

// Warnf : warn msg with format
func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

// Warnln : warn msg + new line
func Warnln(args ...interface{}) {
	globalLogger.Warnln(args...)
}

// Info : info msg
func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

// Infof : info msg with format
func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

// Infoln : info msg + new line
func Infoln(args ...interface{}) {
	globalLogger.Infoln(args...)
}

// Debug : debug msg
func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

// Debugf : debug msg with format
func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

// Debugln : debug msg + new line
func Debugln(args ...interface{}) {
	globalLogger.Debugln(args...)
}

// Trace : trace msg
func Trace(args ...interface{}) {
	globalLogger.Trace(args...)
}

// Tracef : trace msg with format
func Tracef(format string, args ...interface{}) {
	globalLogger.Tracef(format, args...)
}

// Traceln : trace msg + new line
func Traceln(args ...interface{}) {
	globalLogger.Traceln(args...)
}
