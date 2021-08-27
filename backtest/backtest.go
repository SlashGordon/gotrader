package backtest

import (
	"github.com/dirkolbrich/gobacktest"
	"github.com/dirkolbrich/gobacktest/data"
	"github.com/dirkolbrich/gobacktest/strategy"
)

func Run(symbols []string, dataDir string) {
	// initiate a new backtester
	test := gobacktest.New()

	test.SetSymbols(symbols)

	// create a data provider and load the data into the backtest
	data := &data.BarEventFromCSVFile{FileDir: dataDir}
	data.Load(symbols)
	test.SetData(data)

	// choose a strategy
	strategy := strategy.BuyAndHold()

	// create an asset and append it to the strategy
	strategy.SetChildren(gobacktest.NewAsset("SHOP"))

	// load the strategy into the backtest
	test.SetStrategy(strategy)

	// run the backtest
	test.Run()

	// print the results of the test
	test.Stats().PrintResult()
}
