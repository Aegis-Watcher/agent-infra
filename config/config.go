package config

import "time"

type Config struct {
	ServerURL          string
	CollectionInterval time.Duration
}
