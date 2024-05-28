package config

import (
    "time"
)

type Config struct {
    ServerURL string
    Interval  time.Duration
}

func LoadConfig() (*Config, error) {
    // For simplicity, hardcoding the values
    // In a real application, load from environment variables or a config file
    return &Config{
        ServerURL: "http://yourserver.com/report",
        Interval:  60 * time.Second,
    }, nil
}
