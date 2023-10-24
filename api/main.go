package main

import (
	"api/database"
	"api/models"
	"api/router"
	"os"
)

func main() {
	// データベースに接続する。
  database.Connect()
    // 指定されたモデルに基づいてデータベースのテーブルを自動生成または更新する。
	database.DB.AutoMigrate(&models.Host{})

	// Ginの新しいインスタンスを作成する。
	r := router.SetupRouter()

	// 環境変数からポート番号を取得する。指定がなければデフォルトの8080を使用する。
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    // 指定されたポートでサーバーを起動する。
	r.Run(":" + port)
}
