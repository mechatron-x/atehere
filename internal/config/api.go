package config

type Api struct {
	Version    string `json:"version"`
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	StaticRoot string `json:"staticRoot"`
	URL        string `json:"url"`
}
