package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	// GORMを使ってデータベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗しました: %v", err)
	}
	// Userモデルに基づいてテーブルを自動で作成/マイグレーション
	// err = db.AutoMigrate(&models.User{}, &models.UserToken{}, &models.Vote{})
	// if err != nil {
	// 	return nil, fmt.Errorf("テーブルのマイグレーションに失敗しました: %v", err)
	// }

	return db, nil
}
