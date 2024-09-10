package config

var DBConfigInfo DBConfig

type DBConfig struct {
	DataSource string `env:"DATA_SOURCE" envDefault:""`
	//UserPassword string `env:"USER_PW" envDefault:""`

	DBTlsEnabled bool   `env:"DB_TLS_ENABLED" envDefault:""`
	DBTlsCa      string `env:"DB_TLS_CA" envDefault:""`
	DBTlsCert    string `env:"DB_TLS_CERT" envDefault:""`
	DBTlsKey     string `env:"DB_TLS_KEY" envDefault:""`
	DBLogLevel   int    `env:"DB_LOG_LEVEL" envDefault:"3"`
}
