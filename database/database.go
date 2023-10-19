package database

import (
	"fmt"
	"log"
	"os"

	"api/models"

	// GORM (Go Object Relational Mapper) をインポートする。
	"github.com/jinzhu/gorm"
	// MySQL ダイアレクトをインポートする。このインポートにより GORM が MySQL に接続できるようになる。
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// .env ファイルをロードするためのライブラリをインポートする。
	"github.com/joho/godotenv"
)

// DB は GORM によるデータベースの接続を保持するグローバル変数。
var DB *gorm.DB

func Connect() {
	// .env ファイルをロードして環境変数を設定する。
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// .env ファイルまたはシステムの環境変数からデータベースの設定情報を取得する。
	driver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// 上記の設定情報を元にデータベースへの接続URIを構築する。
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// 指定された設定情報でデータベースへ接続する。
	DB, err = gorm.Open(driver, dbURI)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// 指定されたモデルに基づいてデータベースのテーブルを自動生成または更新する。
	DB.AutoMigrate(&models.Host{})
}
