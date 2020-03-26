package main

import (
	// ロガー
	"log"
	"net/http"

	// コントローラー
	"webchat/controllers"

	// Gin セッション情報
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	// Gin Freamwork
	"github.com/gin-gonic/gin"
)

// main エントリーポイント
func main() {

	// セットアップ
	setup()
}

// setup セットアップ
func setup() {

	// Ginルータ作成
	// ※Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェアを保持
	router := gin.Default()

	// session設定
	// ・LoginSession -> ログイン情報
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("LoginSession", store))

	// 静的ファイルのパスを指定
	//router.Static("/assets", "./assets")
	//router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")
	router.StaticFS("/webchat", http.Dir("views"))

	// APIグルーピング
	/*api := router.Group("/api") {
		router.GET("/findAllLoginUsers", controllers.FindAllLoginUsers)
		router.GET("/findLoginUser", controllers.FindLoginUser)
		router.POST("/addLoginUser", controllers.AddLoginUser)
	}*/
	router.GET("/findAllLoginUsers", controllers.FindAllLoginUsers)
	router.GET("/findLoginUser", controllers.FindLoginUser)
	router.POST("/addLoginUser", controllers.AddLoginUser)

	/*
		// ハンドラ設定
		// todo loadtemplatesとして取りまとめたほうがかっこいいかも
		router.GET("/", controllers.Home)
		router.GET("/login", controllers.LogIn)
		router.GET("/signup", controllers.SignUp)
		router.GET("/chat", controllers.Chat)
		router.NoRoute(controllers.NoRoute)
	*/

	// Server 起動
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
