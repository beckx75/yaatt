/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"beckx.online/yaatt/yaatt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "just a test cmd to test writing ID3 Tags",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		yd, err = yaatt.NewYaattData(args, ".")
		if err != nil {
			log.Error().Msgf("%v", err)
		}

		err = yaatt.WriteMetadata(yd.Files[0], *yd.Tagmap)
		if err != nil {
			log.Error().Msgf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
