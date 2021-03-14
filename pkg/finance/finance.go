package finance 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/piquette/finance-go/quote"
	"fmt"
	"unsafe"
	"github.com/FriendlyUser/go_fin_server/pkg/types"
)

// ShowTickers godoc
// @Summary Get Yahoo stock tickers
// @Description get tickers in pandas format
// @Accept  json
// @Produce  json
// @Success 200 {object} Message
// @Failure 400 {object} types.HTTPError
// @Failure 404 {object} types.HTTPError
// @Failure 500 {object} types.HTTPError
// @Router /tickers [get]
func ShowTickers(c *fiber.Ctx) error {
	
	quotes := queryMulti(c, "quotes")
	iter := quote.List(quotes)
	var stock_data [][]string
	var used_symbols []string
	// Iterate over results. Will exit upon any error.
	columns := [8]string{"Last Price", 
	"Change", "Volume", "Avg Vol (3 Month)", "Vol Ratio", 
	"Dollar", "Market", "Exchange"}
	for iter.Next() {
		q := iter.Quote()
		volume_ratio := float64(q.RegularMarketVolume) / float64(q.AverageDailyVolume3Month)
		used_symbols = append(used_symbols, q.Symbol)
		stock_data = append(stock_data, []string{ 
			fmt.Sprintf("%2.2f", q.RegularMarketPrice),   
			fmt.Sprintf("%2.2f", q.RegularMarketChangePercent), 
			fmt.Sprintf("%d", q.RegularMarketVolume),  
			fmt.Sprintf("%d", q.AverageDailyVolume3Month),
			fmt.Sprintf("%2.2f", volume_ratio),
			q.CurrencyID,
			q.MarketID,
			q.ExchangeID})
	}

	return c.JSON(types.Message{Data: stock_data, Index: used_symbols, Columns: columns})
}


func queryMulti(ctx *fiber.Ctx, key string) (values []string) {
	valuesBytes := ctx.Context().QueryArgs().PeekMulti(key)
	values = make([]string, len(valuesBytes))
	for i, v := range valuesBytes {
		values[i] = getString(v)
	}
	return values
}

// #nosec G103
// getString converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
var getString = func(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
