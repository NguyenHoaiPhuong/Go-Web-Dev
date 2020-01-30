package main

import (
	"time"

	"go.uber.org/zap"
)

func testT1()  {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",	// key - value pair
		"attempt", 3,					// key - value pair
		"backoff", time.Second,			// key - value pair
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")	// formatted msg
}

func main() {
	testT1()
}
