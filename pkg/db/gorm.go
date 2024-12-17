package db

import (
	"fmt"
	"github.com/CracherX/favorite_hist/internal/usecase/repository/models"
	"github.com/CracherX/favorite_hist/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

func Connect(cfg *config.Config, retr int) (db *gorm.DB, err error) {
	dsn := parseConfigDSN(cfg)
	for i := 0; i <= retr; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			err = db.AutoMigrate(&models.Favorite{})
			return db, nil
		}
		rem := retr - i
		log.Printf("Ошибка подключения к БД. Попытка... Отсалось попыток: %d", rem)
		time.Sleep(5)
	}
	return nil, err
}

func parseConfigDSN(cfg *config.Config) string {
	params := map[string]string{
		"host":     cfg.Database.Host,
		"port":     cfg.Database.Port,
		"user":     cfg.Database.User,
		"password": cfg.Database.Password,
		"dbname":   cfg.Database.Name,
		"sslmode":  cfg.Database.SslMode,
	}
	var dsnParts []string
	for key, value := range params {
		dsnParts = append(dsnParts, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(dsnParts, " ")
}
