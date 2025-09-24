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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows the metadata-informations of an audiofile",
	Long: `give files or folder-names.

'show' will search for mp3 and flac files.
For each audiofile audio-metadata will be read and printed out`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		yd, err = yaatt.NewYaattData(args, ".")
		if err != nil {
			log.Error().Msgf("%v", err)
		}
		fmt.Printf("%v\n", yd.Tagmap.Id323ToYatt)
		fmt.Println("Got Audiofiles:", len(yd.Files))

		// fmt.Println(yd.PrintMetadata())
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
