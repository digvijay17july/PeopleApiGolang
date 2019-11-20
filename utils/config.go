package utils

type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	PortNo   int
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "password",
			Name:     "testDb",
			Host:     "localhost",
			Charset:  "utf8",
			PortNo:   3306,
		},
	}
}
