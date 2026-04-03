package config

type AppConfigs struct {
	Name    string
	Env     string
	Debug   bool
	Host    string
	Port    string
	Key     string
	ExpMins int64
}

var App *AppConfigs

func init() {
	App = &AppConfigs{
		Name:    GetEnv("APP_NAME", "goxpress"),
		Env:     GetEnv("APP_ENV", "development"),
		Debug:   GetEnvAsBool("APP_DEBUG", false),
		Host:    GetEnv("APP_URL", "127.0.0.1"),
		Port:    GetEnv("APP_PORT", "8080"),
		Key:     GetEnv("JWT_SECRET_KEY", ""),
		ExpMins: GetEnvAsInt("JWT_EXPIRE_MINUTES", 10080),
	}
}
