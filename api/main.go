package main

import (
	"api/controllers"
	"api/middlewares"
	"api/database"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースに接続する。
    database.Connect()

    // Ginの新しいインスタンスを作成する。
	router := gin.Default()

    // 公開APIエンドポイントのグループを作成する。
	public := router.Group("/api")

    // ユーザー登録とログインのルートを設定する。
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

    // 保護されたAPIエンドポイントのグループを作成する。
	protected := router.Group("/api/admin")
	// JWT認証ミドルウェアを保護されたエンドポイントに適用する。
	protected.Use(middlewares.JwtAuthMiddleware())
	// 現在のユーザー情報を取得するルートを設定する。
	protected.GET("/user", controllers.CurrentUser)

	// 環境変数からポート番号を取得する。指定がなければデフォルトの8080を使用する。
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    // 指定されたポートでサーバーを起動する。
	router.Run(":" + port)
}
