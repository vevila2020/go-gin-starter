package database

import (
	"fmt"

	"go_basic_framework_template/internal/config"
	"go_basic_framework_template/internal/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("config not initialized")
	}

	dbCfg := cfg.Database

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
