package config

type Config struct {
	IsDebug *bool `yaml: "is_debug"`
	Listen struct {
		Type string `yaml: "type"`
		BindIP string `yaml: "bind_ip"`
		Port string `yaml: "port"`
	} `yaml:"listen"` 
}