package config

import (
	"sync"

	"github.com/Freshify-HHS-ICT/backend/structs"
	"github.com/go-playground/validator/v10"
)

var (
	configInstance *structs.Config
	configOnce     sync.Once
	Validator      = validator.New()
)

func GetConfig() *structs.Config {
	configOnce.Do(func() {
		configInstance = &structs.Config{
			Server: &structs.ServerConfig{
				LogLevel: getEnvAsString("LOG_LEVEL", "INFO"),
			},
			Database:  &structs.DatabaseConfig{},
			Cors:      &structs.CorsConfig{},
			Cache:     &structs.CacheConfig{},
			RateLimit: &structs.RateLimitConfig{},
			Auth:      &structs.AuthConfig{},
		}
	})

	return configInstance
}
