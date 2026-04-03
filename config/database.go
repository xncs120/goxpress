package config

type DBConfigs struct {
	Url string
}

var Database *DBConfigs

func init() {
	Database = &DBConfigs{
		Url: GetEnv("DB_URL", ""),
	}
}
