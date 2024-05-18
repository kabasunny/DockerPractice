package main

import (
	"net/http"          // HTTPクライアントとサーバーの機能を提供
	"net/http/httptest" // HTTPリクエストとレスポンスのテストをサポート
	"testing"           // テストパッケージ

	"github.com/gin-gonic/gin"           // Ginフレームワーク
	"github.com/stretchr/testify/assert" // assertパッケージは、テストのアサーション（期待値と実際の値の比較）を提供
)

// TestPingRoute関数は、/pingエンドポイントのテストを行う
func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)  // テストモードを設定。これにより、ログ出力が抑制。

	// デフォルトのGinルーターを作成。
	r := gin.Default()
	// /pingエンドポイントにGETリクエストが来たときのハンドラーを設定。
	r.GET("/ping", func(c *gin.Context) {
		// ステータスコード200とともにJSONレスポンスを返す。
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// /pingエンドポイントへのGETリクエストを作成。
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	// レスポンスを記録するためのレコーダーを作成。
	resp := httptest.NewRecorder()
	// リクエストをルーターに送信し、レスポンスをレコーダーに記録。
	r.ServeHTTP(resp, req)

	// レスポンスのステータスコードが200であることを確認。
	assert.Equal(t, http.StatusOK, resp.Code)
	// レスポンスボディが期待通りであることを確認。
	assert.Equal(t, `{"message":"pong"}`, resp.Body.String())
}
