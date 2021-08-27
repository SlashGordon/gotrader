package cmd

import (
	"github.com/SlashGordon/gotrader/trader"
	"github.com/SlashGordon/gotrader/utils"
	"github.com/spf13/cobra"
)

func init() {

	var (
		symbols []string
		dataDir string
		days    int
	)

	startCmd := &cobra.Command{
		Use:   "dump",
		Short: "downloads price infos for given stocks",
		Run: func(cmd *cobra.Command, args []string) {
			client := trader.NewTrader()
			for _, sym := range symbols {
				utils.Logger.Info("Dump data for " + sym)
				if err := client.WriteBarToCSV(sym, days, dataDir); err != nil {
					utils.Logger.Error(err)
				}
			}
		},
	}
	configFlag := startCmd.PersistentFlags()
	configFlag.StringSliceVarP(&symbols, "symbols", "s", nil, "list of symbols")
	configFlag.StringVarP(&dataDir, "data", "d", "data", "data dir")
	configFlag.IntVarP(&days, "days", "l", 300, "Set history data size")
	cobra.MarkFlagRequired(configFlag, "config")

	RootCmd.AddCommand(startCmd)
}
