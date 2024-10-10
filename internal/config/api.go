package config

type Api struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	WebRoot string `json:"webRoot"`
}
