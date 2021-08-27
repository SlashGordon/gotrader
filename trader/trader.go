package trader

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	v2 "github.com/alpacahq/alpaca-trade-api-go/v2"
)

type trader struct {
	client *alpaca.Client
	acc    *alpaca.Account
}

func NewTrader() *trader {
	alpacaClient := alpaca.NewClient(common.Credentials())
	acc, _ := alpacaClient.GetAccount()
	return &trader{
		client: alpacaClient,
		acc:    acc,
	}
}

func (t *trader) StockPriceDailyClose(stock string, lookBack int) []float64 {
	barset := make([]float64, 0, lookBack)
	bars := t.client.GetBars(
		stock, v2.Day, v2.Raw, time.Now().Add(-time.Duration(lookBack)*24*time.Hour), time.Now().Add(-20*time.Minute), lookBack)

	for bar := range bars {
		if bar.Error != nil {
			panic(bar.Error)
		}
		barset = append(barset, bar.Bar.Close)
	}

	return barset
}

func (t *trader) StockPriceDaily(stock string, lookBack int) []v2.Bar {
	barset := make([]v2.Bar, 0, lookBack)
	bars := t.client.GetBars(
		stock, v2.Day, v2.Raw, time.Now().Add(-time.Duration(lookBack)*24*time.Hour), time.Now().Add(-20*time.Minute), lookBack)

	for bar := range bars {
		if bar.Error != nil {
			panic(bar.Error)
		}
		barset = append(barset, bar.Bar)
	}

	return barset
}

func (t *trader) WeeklyStockPrice(stock string) []v2.Bar {

	bars := t.client.GetBars(
		stock, v2.Day, v2.Raw, time.Now().Add(-7*24*time.Hour), time.Now().Add(-20*time.Minute), 7)
	var barset []v2.Bar

	for bar := range bars {
		if bar.Error != nil {
			panic(bar.Error)
		}
		barset = append(barset, bar.Bar)
	}

	return barset
}
