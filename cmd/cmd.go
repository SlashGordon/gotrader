package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/SlashGordon/gotrader/utils"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd global entry point
var RootCmd = &cobra.Command{
	Use: "gotrader",
}

// Execute ...
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotrader.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gotrader")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		utils.Logger.Info("Using config file:", viper.ConfigFileUsed())
		// Check for environment variables
		if common.Credentials().ID == "" {
			os.Setenv(common.EnvApiKeyID, viper.GetString(common.EnvApiKeyID))
		}
		if common.Credentials().Secret == "" {
			os.Setenv(common.EnvApiSecretKey, viper.GetString(common.EnvApiSecretKey))
		}
		alpaca.SetBaseUrl(viper.GetString("APCA_API"))
		utils.Logger.Infof("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)
	}
}
