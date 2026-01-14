package main

import (
	"moon/pkg/config"
	"moon/pkg/log"
)

func main() {
	cfg := &config.Configuration{}
	// 数据库配置
	cfg.DatabaseConfiguration.Driver = "postgres"
	cfg.DatabaseConfiguration.DSN = "postgres://capstone:capstone@localhost:5432/capstone_test"
	cfg.DatabaseConfiguration.MaxOpenConnections = 50
	cfg.DatabaseConfiguration.MaxIdleConnections = 10

	// JWT配置
	cfg.JWTConfiguration.Secret = []byte("test-secret-key-min-32-bytes-long!!!!")
	cfg.JWTConfiguration.AccessTokenTTL = 15
	cfg.JWTConfiguration.RefreshTokenTTL = 7
	cfg.JWTConfiguration.OverlapWindow = 5
	cfg.JWTConfiguration.Issuer = "test-app"
	cfg.LogConfiguration.Level = 1
	cfg.LogConfiguration.Format = "json"

	log := log.New(cfg.LogConfiguration.Level, cfg.LogConfiguration.Format)

	log.Info("Configuration Initialized")
}
