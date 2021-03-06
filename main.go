package main

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/FriendlyUser/go_fin_server/pkg/finance"
	"github.com/FriendlyUser/go_fin_server/pkg/rssData"
	"github.com/FriendlyUser/go_fin_server/pkg/types"
	_ "github.com/FriendlyUser/go_fin_server/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// TODO function to validate env vars goes here
	app := fiber.New()
	// view swagger docs at http://localhost:8080/swagger/index.html
	app.Get("/swagger/*", fiberSwagger.Handler)
	
	app.Get("/tickers", finance.ShowTickers)

	app.Get("/rss-data", rssData.GetRssData)
	app.Get("/accounts/:id", ShowAccount)
	app.Listen(":8080")
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} types.Account
// @Failure 400 {object} types.HTTPError
// @Failure 404 {object} types.HTTPError
// @Failure 500 {object} types.HTTPError
// @Router /accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	return c.JSON(types.Account{
		Id: c.Params("id"),
	})
}
