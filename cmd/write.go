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

		tts, ok := yd.MetaDatas[yd.Files[0]].TextTags["Title"]
		if ok {
			fmt.Println(tts[0].Value)
			tts[0].Value = "Reverse"
		} else {
			yd.MetaDatas[yd.Files[0]].TextTags["Title"] = []*yaatt.TextTag{
				{OrgName: "TIT2", Value: "Sepps Song"},
			}
		}
		tts, ok = yd.MetaDatas[yd.Files[0]].TextTags["Album"]
		if ok {
			fmt.Println(tts[0].Value)
			tts[0].Value = "Album Reverse"
		} else {
			yd.MetaDatas[yd.Files[0]].TextTags["Album"] = []*yaatt.TextTag{
				{OrgName: "TALB", Value: "An Sepp sei Album"},
			}
		}
		for fp, md := range yd.MetaDatas {
			err = md.WriteMetadata(fp, *yd.Tagmap)
			if err != nil {
				log.Error().Msgf("%v", err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
