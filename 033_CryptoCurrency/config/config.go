package config

import (
	"log"
	"os"
)

var config *Config

func init() {
	// cmc
	cmcAPIKey := os.Getenv("cmc_api_key")

	// blockcypher
	blockCypherAPIKey := os.Getenv("blockcypher_api_token")
	blockCypherBTCEndpoint := os.Getenv("blockcypher_btc_endpoint")

	// we don't use zap log here for easy development
	log.Printf(`
		cmcAPIKey: %s
		blockCypherAPIKey: %s
		blockCypherBTCEndpoint: %s
		`,
		cmcAPIKey,
		blockCypherAPIKey,
		blockCypherBTCEndpoint,
	)

	config = &Config{
		CMCApiKey:              cmcAPIKey,
		BlockCypherAPIKey:      blockCypherAPIKey,
		BlockCypherBTCEndpoint: blockCypherBTCEndpoint,
	}
}

// GetConfig :
func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	// CMC
	CMCApiKey string `json:"cmc_api_key"`

	// Blockcypher
	BlockCypherAPIKey      string `json:"blockcypher_api_token"`
	BlockCypherBTCEndpoint string `json:"blockcypher_btc_endpoint"`
}
