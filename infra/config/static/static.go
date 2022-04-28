package static

import (
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	MySQL struct {
		DSN             string        `toml:"dsn"`
		MaxIdleConns    int           `toml:"max_idle_conns"`
		MaxOpenConns    int           `toml:"max_open_conns"`
		MaxConnLifeTime time.Duration `toml:"max_conn_life_time"`
	} `toml:"mysql""`
}

// ReadConfig read static config
func ReadConfig() (Config, error) {
	var c Config
	_, err := toml.DecodeFile("./config.toml", &c)
	return c, err
}
