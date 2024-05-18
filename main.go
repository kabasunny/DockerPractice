package main

import (
	"github.com/gin-gonic/gin" // Ginフレームワークをインポート
)

func main() {  
	r := gin.Default()  // デフォルトのGinルーターを作成

	// "/ping"エンドポイントにGETリクエストが来たときのハンドラーを設定
	r.GET("/ping", func(c *gin.Context) {
		// ステータスコード200とともにJSONレスポンスを返す
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()  // デフォルトのHTTPアドレス（:8080）でサーバーを起動
}
