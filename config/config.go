package config

type (
	// Config .
	Config struct {
		// Base
		Base struct {
			Port string `yaml:"port"`
		} `yaml:"base"`

		// Postgresql.
		Postgresql struct {
			Url string `yaml:"url"`
		} `ymal:"postgresql"`
	}
)
