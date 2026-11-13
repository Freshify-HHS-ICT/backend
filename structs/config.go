package structs

type Config struct {
	Server    *ServerConfig
	Database  *DatabaseConfig
	Cors      *CorsConfig
	Cache     *CacheConfig
	Auth      *AuthConfig
	RateLimit *RateLimitConfig
}

type ServerConfig struct {
	LogLevel string
}

type DatabaseConfig struct {
}

type CorsConfig struct {
}

type AuthConfig struct {
}

type CacheConfig struct {
}

type RateLimitConfig struct {
}
