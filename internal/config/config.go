package config

type Config struct {
	ApiKey        string `env:"API_KEY" yaml:"api_key"`
	ApiSecret     string `env:"API_SECRET" yaml:"api_secret"`
	ApiPassphrase string `env:"API_PASSPHRASE" yaml:"api_passphrase"`
}
