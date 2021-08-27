package cmd

import (
	"github.com/SlashGordon/gotrader/backtest"
	"github.com/spf13/cobra"
)

func init() {

	var (
		strategyName  string
		symbols       []string
		dataDir       string
		days          int
		portfolioCash float64
	)

	startCmd := &cobra.Command{
		Use:   "backtest",
		Short: "Start backtest for given strategy",
		Run: func(cmd *cobra.Command, args []string) {
			backtest.Run(symbols, dataDir)
		},
	}
	configFlag := startCmd.PersistentFlags()
	configFlag.StringVarP(&strategyName, "strategy", "s", "", "name of strategy")
	configFlag.StringVarP(&dataDir, "data", "d", "data", "data dir")
	configFlag.IntVarP(&days, "days", "l", 300, "lookback in days. If lookback is set to 300 the backtest runs from now -300 until now.")
	configFlag.Float64VarP(&portfolioCash, "portfolioCash", "c", 1000.0, "Initial cash of the porfolio.")
	configFlag.StringSliceVarP(&symbols, "symbols", "y", nil, "list of symbols")
	cobra.MarkFlagRequired(configFlag, "config")

	RootCmd.AddCommand(startCmd)
}
