/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"beckx.online/yaatt/yaatt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "read in given files and sniff on them",
	Long: `read in given file/folders

make some collections like "genres", etc`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		yd, err = yaatt.NewYaattData(args, ".")
		if err != nil {
			log.Error().Msgf("%v", err)
			return
		}
		fmt.Printf("%v\n", yd.Tagmap.Id323ToYatt)
		fmt.Println("Got Audiofiles:", len(yd.Files))
	},
}

func init() {
	rootCmd.AddCommand(sniffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sniffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sniffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
