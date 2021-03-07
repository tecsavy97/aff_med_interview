package models

var AppConfig Config

// User - Object for User for username and password
type User struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

// AppConfig - Object for App configurations
type Config struct {
	Port   string
	JwtKey string
}

func (a *Config) GetConfigFilePath() string {
	return "config.toml"
}
