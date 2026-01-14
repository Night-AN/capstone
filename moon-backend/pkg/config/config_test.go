package config

import (
	"encoding/json"
	"os"
	"sync"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// testConfig 定义测试用的配置数据
var testConfig = &Configuration{
	DatabaseConfiguration: DatabaseConfiguration{
		Driver:             "postgres",
		DSN:                "host=localhost user=capstone password=capstone dbname=capstone_test port=5432 sslmode=disable",
		MaxOpenConnections: 50,
		MaxIdleConnections: 5,
	},
	JWTConfiguration: JWTConfiguration{
		Secret:          []byte("test-secret-key-12345"),
		AccessTokenTTL:  60,   // 60分钟
		RefreshTokenTTL: 1440, // 24小时（24*60）
		OverlapWindow:   5,    // 5分钟
		Issuer:          "test-service",
	},
	LogConfiguration: LogConfiguration{
		Level:  2,
		Format: "json",
	},
}

// getTestDSN 获取测试数据库连接字符串
func getTestDSN() string {
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=capstone password=capstone dbname=capstone_test port=5432 sslmode=disable"
	}
	return dsn
}

// setupTestDB 初始化测试数据库
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := getTestDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// 确保使用与主包相同的表名
	if err := db.AutoMigrate(&SystemConfig{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	// 清理测试数据
	db.Where("service_name LIKE ?", "test-%").Delete(&SystemConfig{})

	return db
}

// teardownTestDB 清理测试数据库
func teardownTestDB(db *gorm.DB) {
	db.Where("service_name LIKE ?", "test-%").Delete(&SystemConfig{})
}

// TestSaveAndLoadConfig 测试配置保存和加载完整流程
func TestSaveAndLoadConfig(t *testing.T) {
	dsn := getTestDSN()
	serviceName := "test-user-service"

	db := setupTestDB(t)
	defer teardownTestDB(db)

	// 保存配置
	if err := SaveToDB(dsn, serviceName, testConfig); err != nil {
		t.Fatalf("SaveToDB failed: %v", err)
	}

	// 加载配置
	loadedConfig, err := LoadFromDB(dsn, serviceName)
	if err != nil {
		t.Fatalf("LoadFromDB failed: %v", err)
	}
	if loadedConfig == nil {
		t.Fatal("Loaded config should not be nil")
	}

	// 验证配置内容
	if loadedConfig.DatabaseConfiguration.Driver != testConfig.DatabaseConfiguration.Driver {
		t.Errorf("Driver mismatch: got %s, want %s",
			loadedConfig.DatabaseConfiguration.Driver, testConfig.DatabaseConfiguration.Driver)
	}
	if loadedConfig.JWTConfiguration.Issuer != testConfig.JWTConfiguration.Issuer {
		t.Errorf("Issuer mismatch: got %s, want %s",
			loadedConfig.JWTConfiguration.Issuer, testConfig.JWTConfiguration.Issuer)
	}
	if !bytesEqual(loadedConfig.JWTConfiguration.Secret, testConfig.JWTConfiguration.Secret) {
		t.Error("Secret mismatch")
	}
}

// TestLoadNonExistentConfig 测试加载不存在的配置
func TestLoadNonExistentConfig(t *testing.T) {
	dsn := getTestDSN()
	serviceName := "test-non-existent-service"

	config, err := LoadFromDB(dsn, serviceName)
	if err == nil {
		t.Error("Should return error for non-existent config")
	}
	if config != nil {
		t.Error("Config should be nil when not found")
	}
}

// TestConfigVersionIncrement 测试配置版本号递增
func TestConfigVersionIncrement(t *testing.T) {
	dsn := getTestDSN()
	serviceName := "test-version-service"

	db := setupTestDB(t)
	defer teardownTestDB(db)

	// 保存初始配置
	if err := SaveToDB(dsn, serviceName, testConfig); err != nil {
		t.Fatalf("First SaveToDB failed: %v", err)
	}

	// 获取第一次保存的版本
	var firstConfig SystemConfig
	db.Where("service_name = ?", serviceName).First(&firstConfig)
	firstVersion := firstConfig.Version

	// 修改配置并再次保存
	testConfig.LogConfiguration.Level = 3
	if err := SaveToDB(dsn, serviceName, testConfig); err != nil {
		t.Fatalf("Second SaveToDB failed: %v", err)
	}

	// 获取第二次保存的版本
	var secondConfig SystemConfig
	db.Where("service_name = ?", serviceName).First(&secondConfig)

	// 验证版本号递增
	if secondConfig.Version != firstVersion+1 {
		t.Errorf("Version should increment: got %d, want %d",
			secondConfig.Version, firstVersion+1)
	}
}

// TestLoadConfigWithInvalidData 测试加载无效数据
func TestLoadConfigWithInvalidData(t *testing.T) {
	dsn := getTestDSN()
	serviceName := "test-invalid-data"

	db := setupTestDB(t)
	defer teardownTestDB(db)

	// 直接插入无效数据
	invalidData := []byte("invalid json data")
	db.Create(&SystemConfig{
		ServiceName: serviceName,
		Config:      invalidData,
		Version:     1,
	})

	// 尝试加载无效配置
	config, err := LoadFromDB(dsn, serviceName)
	if err == nil {
		t.Error("Should return error for invalid data")
	}
	if config != nil {
		t.Error("Config should be nil when data is invalid")
	}
}

// TestJSONSerialization 测试JSON序列化完整性
func TestJSONSerialization(t *testing.T) {
	// 序列化
	configBytes, err := json.Marshal(testConfig)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}

	// 反序列化
	var decodedConfig Configuration
	if err := json.Unmarshal(configBytes, &decodedConfig); err != nil {
		t.Fatalf("JSON unmarshal failed: %v", err)
	}

	// 验证
	if decodedConfig.JWTConfiguration.AccessTokenTTL != testConfig.JWTConfiguration.AccessTokenTTL {
		t.Errorf("AccessTokenTTL mismatch after JSON round trip")
	}
}

// TestConcurrentConfigAccess 测试并发访问安全性
func TestConcurrentConfigAccess(t *testing.T) {
	dsn := getTestDSN()
	serviceName := "test-concurrent-service"

	db := setupTestDB(t)
	defer teardownTestDB(db)

	if err := SaveToDB(dsn, serviceName, testConfig); err != nil {
		t.Fatalf("SaveToDB failed: %v", err)
	}

	done := make(chan bool)
	errCount := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			config, err := LoadFromDB(dsn, serviceName)
			if err != nil || config == nil {
				mu.Lock()
				errCount++
				mu.Unlock()
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if errCount > 0 {
		t.Errorf("Concurrent access had %d failures", errCount)
	}
}

// BenchmarkSaveToDB 基准测试：保存配置
func BenchmarkSaveToDB(b *testing.B) {
	dsn := getTestDSN()
	serviceName := "test-benchmark-save"

	// 基准测试也需要清理
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	defer db.Where("service_name = ?", serviceName).Delete(&SystemConfig{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := SaveToDB(dsn, serviceName, testConfig); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLoadFromDB 基准测试：加载配置
func BenchmarkLoadFromDB(b *testing.B) {
	dsn := getTestDSN()
	serviceName := "test-benchmark-load"

	SaveToDB(dsn, serviceName, testConfig)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := LoadFromDB(dsn, serviceName); err != nil {
			b.Fatal(err)
		}
	}
}

// bytesEqual 辅助函数：比较两个字节数组
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
