package config

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SystemConfig 系统配置模型 - 存储在数据库中
type SystemConfig struct {
	ServiceName string    `gorm:"column:service_name;type:varchar(255);primaryKey;comment:服务名称" mapstructure:"service_name"`
	Config      []byte    `gorm:"column:config;type:bytea;not null;comment:配置二进制数据" mapstructure:"config"`
	Version     int       `gorm:"column:version;type:integer;default:1;comment:配置版本" mapstructure:"version"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;comment:最后更新时间" mapstructure:"updated_at"`
}

// TableName 指定数据库表名
func (SystemConfig) TableName() string {
	return "system_configs"
}

// Configuration 运行时配置结构
type Configuration struct {
	DatabaseConfiguration DatabaseConfiguration `mapstructure:"database" json:"database"`
	JWTConfiguration      JWTConfiguration      `mapstructure:"jwt" json:"jwt"`
	LogConfiguration      LogConfiguration      `mapstructure:"log" json:"log"`
}

// DatabaseConfiguration 数据库配置
type DatabaseConfiguration struct {
	Driver             string `mapstructure:"driver" json:"driver"`
	DSN                string `mapstructure:"dsn" json:"dsn"`
	MaxOpenConnections int    `mapstructure:"max_open_connections" json:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections" json:"max_idle_connections"`
}

// JWTConfiguration JWT相关配置
type JWTConfiguration struct {
	Secret          []byte `mapstructure:"secret" json:"secret"`
	AccessTokenTTL  int    `mapstructure:"access_token_ttl" json:"access_token_ttl"`
	RefreshTokenTTL int    `mapstructure:"refresh_token_ttl" json:"refresh_token_ttl"`
	OverlapWindow   int    `mapstructure:"overlap_window" json:"overlap_window"`
	Issuer          string `mapstructure:"issuer" json:"issuer"`
}

// LogConfiguration 日志配置
type LogConfiguration struct {
	Level  int    `mapstructure:"level" json:"level"`
	Format string `mapstructure:"format" json:"format"`
}

// LoadFromDB 从PostgreSQL数据库加载配置
func LoadFromDB(configDsn, serviceName string) (*Configuration, error) {
	// 连接配置数据库
	db, err := gorm.Open(postgres.Open(configDsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接配置数据库失败: %w", err)
	}

	// 查询配置
	var sysConfig SystemConfig
	err = db.Where("service_name = ?", serviceName).First(&sysConfig).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("未找到服务配置: %s", serviceName)
		}
		return nil, fmt.Errorf("查询配置失败: %w", err)
	}

	// 反序列化JSON
	var config Configuration
	if err := json.Unmarshal(sysConfig.Config, &config); err != nil {
		return nil, fmt.Errorf("反序列化配置失败: %w", err)
	}

	return &config, nil
}

// SaveToDB 保存配置到数据库
func SaveToDB(configDsn, serviceName string, config *Configuration) error {
	// 连接数据库
	db, err := gorm.Open(postgres.Open(configDsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 序列化为JSON
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}

	// 创建或更新
	return db.Transaction(func(tx *gorm.DB) error {
		var existing SystemConfig
		err := tx.Where("service_name = ?", serviceName).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			return tx.Create(&SystemConfig{
				ServiceName: serviceName,
				Config:      configBytes,
				Version:     1,
			}).Error
		} else if err != nil {
			return fmt.Errorf("查询现有配置失败: %w", err)
		}

		existing.Config = configBytes
		existing.Version += 1
		return tx.Save(&existing).Error
	})
}
