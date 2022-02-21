package main

import (
	_ "ginswag/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title APIドキュメントのタイトル
// @version バージョン(1.0)
// @description 仕様書に関する内容説明
// @termsOfService 仕様書使用する際の注意事項

// @contact.name APIサポーター
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.New()
	swagUrl := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagUrl))

	r.GET("/", test)
	r.Run()
}

// @description テスト用APIの詳細
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "必須ではありません。"
// @Success 200 {object} map[string]interface{} {"code":200,"hoge":"moge"}
// @router / [get]
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hoge": "moge"})
}
