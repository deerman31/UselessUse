package routing

import (
	"backendUselessUse/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/* Routingを表す
 */
func NewRouter() *echo.Echo {
	e := echo.New()

	/* e.Use(middleware.Recover())は、RecoverミドルウェアをEchoインスタンスに追加します。
	   このミドルウェアは、パニックが発生した場合にアプリケーションをクラッシュさせずにリカバリするためのものです。 */
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", handlers.HelloWorldHandler)

	// APIグループの作成
	api := e.Group("/api")
	api.GET("/", handlers.HelloApiHandler)

	/* ここに各apiのエンドポイントを書いていく */
	return e
}
