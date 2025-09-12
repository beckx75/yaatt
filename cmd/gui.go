/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"beckx.online/yaatt/gui"
	"github.com/spf13/cobra"
)

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "starts the graphical-user-interface",
	Long:  `nothing more to say...`,
	Run: func(cmd *cobra.Command, args []string) {
		gui.InitGui(args)
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
