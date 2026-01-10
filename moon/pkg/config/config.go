package config

type Configuration struct {
	ApplicationName       string
	DatabaseConfiguration struct {
		Driver             string
		Dsn                string
		MaxOpenConnections int
		MaxIdleConnections int
	}
}
