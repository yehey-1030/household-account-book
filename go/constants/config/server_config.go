package config

var ServerConfigInfo ServerConfig

type ServerConfig struct {
	UrlAccessCheckEnabled bool     `env:"URL_ACCESS_CHECK_ENABLED" envDefault:"false"`
	AccessGrantedUrls     []string `env:"ACCESS_GRANTED_URLS" envSeparator:" " envDefault:""`
}
