package main

import (
	"time"

	"go.uber.org/zap"
)

// sugar example
func testT1() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com", // key - value pair
		"attempt", 3, // key - value pair
		"backoff", time.Second, // key - value pair
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com") // formatted msg
}

// logger example
func testT2() {
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// Converting logger and sugar
func testT3() {
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("It is logger message",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow("It is sugar message",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}

func main() {
	// testT1()
	// testT2()
	testT3()
}
