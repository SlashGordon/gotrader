package trader

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func (t *trader) WriteBarToCSV(symbol string, lookback int, datadir string) error {
	loc, _ := time.LoadLocation("America/New_York")
	head := []string{"Date", "Open", "High", "Low", "Close", "Adj Close", "Volume"}

	bars := t.StockPriceDaily(symbol, lookback)

	// create output file for stock
	outputPath := filepath.Join(datadir, symbol+".csv")
	if _, err := os.Stat(datadir); os.IsNotExist(err) {
		err = os.MkdirAll(datadir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(head)
	if err != nil {
		return err
	}

	for _, bar := range bars {
		value := []string{
			bar.Timestamp.In(loc).Format(time.RFC3339)[:10],
			fmt.Sprintf("%f", bar.Open),
			fmt.Sprintf("%f", bar.High),
			fmt.Sprintf("%f", bar.Low),
			fmt.Sprintf("%f", bar.Close),
			fmt.Sprintf("%f", bar.Close),
			fmt.Sprintf("%v", bar.Volume),
		}
		err := writer.Write(value)
		if err != nil {
			return err
		}

	}
	return nil
}
