package config

type DB struct {
	Driver   string `json:"driver"`
	Name     string `json:"Name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Timeout  string `json:"timeout"`
	TryCount uint   `json:"tryCount"`
}
