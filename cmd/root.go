/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"beckx.online/yaatt/yaatt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var yd *yaatt.YaattData

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yaatt",
	Short: "yat again another tag tool",
	Long: `yaatt - my seventh attemp for my own tool to tag audiofiles...

hope this one is my last attemp ;)
`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	userhome, _ := os.UserHomeDir()
	viper.AddConfigPath(".")
	viper.AddConfigPath(userhome + "/yaatt/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Msgf("%v", err)
	}
	switch viper.Get("App.loglevel") {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	log.Info().Msg("initialised root command... starting yaatt")
	log.Debug().Msg("Is this debugging!?")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
