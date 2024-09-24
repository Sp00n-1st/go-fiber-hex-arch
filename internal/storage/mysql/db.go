package mysql

import (
	"go-fiber-hex-arch/internal/adapter/config"
	"go-fiber-hex-arch/internal/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

func ConnectMySQL(cfg *config.DB) (*gorm.DB, error) {
	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	for i := 0; i < 3; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Error("Failed to open MySQL connection", "error", err)
			time.Sleep(2 * time.Second)
			continue
		}

		_, err := db.DB()
		if err != nil {
			slog.Error("Failed to ping MySQL", "error", err)
			time.Sleep(2 * time.Second)
			continue
		}

		slog.Info("Connected MySQL")

		err = db.AutoMigrate(&domain.Product{})
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return nil, err
}
