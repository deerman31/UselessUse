package routing

import (
	"net/http"

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// APIグループの作成
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, api!")
	})
	/* ここに各apiのエンドポイントを書いていく */
	return e
}