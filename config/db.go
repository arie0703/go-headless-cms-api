package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// .envファイルを読み込む（無視しても環境変数から取得可能）
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env ファイルは読み込まれませんでした")
	}

	// 環境変数から取得
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	// DSN（Data Source Name）作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	var db *gorm.DB
	var err error

	maxRetries := 10
	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil && sqlDB.Ping() == nil {
				log.Println("✅ DB接続成功")
				return db
			}
		}
		log.Printf("DB接続リトライ中 (%d/%d): %v", i, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("DB接続失敗: %v", err)
	return nil
}
