package database

import (
	"errors"
	"fmt"

	"go-gin-starter/internal/config"
	"go-gin-starter/internal/logger"
	"go-gin-starter/internal/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

var ErrDisabled = errors.New("database is disabled")

func Init() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("config not initialized")
	}

	dbCfg := cfg.Database

	if !dbCfg.Enabled {
		logger.Info("database is disabled, skipping initialization")
		return nil
	}

	var dialector gorm.Dialector
	switch dbCfg.Driver {
	case "postgres":
		dialector = postgres.Open(dbCfg.DSN())
	case "mysql":
		dialector = mysql.Open(dbCfg.DSN())
	default:
		return fmt.Errorf("unsupported database driver: %s", dbCfg.Driver)
	}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)

	if dbCfg.AutoMigrate {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			return fmt.Errorf("failed to auto migrate: %w", err)
		}
	}

	DB = db

	logger.Info("database connected successfully",
		zap.String("driver", dbCfg.Driver),
		zap.String("host", dbCfg.Host),
		zap.String("database", dbCfg.Database),
	)

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func IsEnabled() bool {
	return DB != nil
}

func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}