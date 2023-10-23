package main

import (
	"api/controllers"
	"api/database"
	"api/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースに接続する。
  database.Connect()
    // 指定されたモデルに基づいてデータベースのテーブルを自動生成または更新する。
	database.DB.AutoMigrate(&models.Host{})

    // Ginの新しいインスタンスを作成する。
	router := gin.Default()

    // 公開APIエンドポイントのグループを作成する。
	public := router.Group("/api")

	public.POST("/hosts", controllers.RegisterHost)
	public.POST("/hosts/login", controllers.LoginHost)

	// 環境変数からポート番号を取得する。指定がなければデフォルトの8080を使用する。
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    // 指定されたポートでサーバーを起動する。
	router.Run(":" + port)
}
