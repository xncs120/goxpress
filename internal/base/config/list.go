package config

type AppConfigs struct {
	Name  string
	Env   string
	Key   string
	Debug bool
	Url   string
	Port  string
}

var App *AppConfigs

func NewAppConfigs() *AppConfigs {
	return &AppConfigs{
		Name:  getEnv("APP_NAME", "goxpress"),
		Env:   getEnv("APP_ENV", "development"),
		Key:   getEnv("APP_KEY", ""),
		Debug: getEnvAsBool("APP_DEBUG", false),
		Url:   getEnv("APP_URL", "http://localhost"),
		Port:  getEnv("APP_PORT", "8080"),
	}
}

type SecretConfigs struct {
	JwtSecret    string
	JwtExpireHrs int64
}

var Secret *SecretConfigs

func NewSecretConfigs() *SecretConfigs {
	return &SecretConfigs{
		JwtSecret:    getEnv("JWT_SECRET", ""),
		JwtExpireHrs: getEnvAsInt("JWT_EXPIRE_HRS", 24),
	}
}

type DBConfigs struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var DB *DBConfigs

func NewDBConfigs() *DBConfigs {
	return &DBConfigs{
		Driver:   getEnv("DB_DRIVER", "postgres"),
		Host:     getEnv("DB_HOST", "db"),
		Port:     getEnv("DB_PORT", "5432"),
		Database: getEnv("DB_DATABASE", "database"),
		Username: getEnv("DB_USERNAME", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
	}
}
