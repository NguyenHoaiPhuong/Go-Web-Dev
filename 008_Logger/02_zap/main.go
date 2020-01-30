// Refer to below link for more details and examples
// https://github.com/uber-go/zap/blob/master/example_test.go
// https://godoc.org/go.uber.org/zap

package main

import (
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func examplePresets() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	logger := zap.NewExample() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	// Output:
	// {"level":"info","msg":"Failed to fetch URL.","url":"http://example.com","attempt":3,"backoff":"1s"}
	// {"level":"info","msg":"Failed to fetch URL: http://example.com"}
	// {"level":"info","msg":"Failed to fetch URL.","url":"http://example.com","attempt":3,"backoff":"1s"}
}

func exampleNamespace() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.With(
		zap.Namespace("metrics"),
		zap.String("firstName", "Nguyen"),
		zap.String("secondName", "Akagi"),
		zap.Int("age", 34),
	).Info("tracked some metrics")
	// Output:
	// {"level":"info","msg":"tracked some metrics","metrics":{"firstName":"Nguyen","secondName":"Akagi","age":34}}
}

func exampleNewStdLog() {
	logger := zap.NewExample()
	defer logger.Sync()

	std := zap.NewStdLog(logger)
	std.Print("standard logger wrapper")
	// Output:
	// {"level":"info","msg":"standard logger wrapper"}
}

func exampleRedirectStdLog() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()

	log.Print("redirected standard library")
	// Output:
	// {"level":"info","msg":"redirected standard library"}
}

func exampleAtomicLevel() {
	atom := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	logger.Info("info logging enabled")

	atom.SetLevel(zap.ErrorLevel)
	logger.Info("info logging disabled")
	// Output:
	// {"level":"info","msg":"info logging enabled"}
}

func exampleFileWriter() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/log.txt",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	writerSyncer := zapcore.AddSync(lumberJackLogger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Starting ...")

	logger.Debug("Finished ...")
}

func main() {
	// examplePresets()
	// exampleNamespace()
	// exampleNewStdLog()
	// exampleRedirectStdLog()
	// exampleAtomicLevel()

	exampleFileWriter()
}
