package config

type DB struct {
	Driver   string `json:"driver"`
	DSN      string `json:"dsn"`
	Timeout  string `json:"timeout"`
	TryCount uint   `json:"tryCount"`
}
