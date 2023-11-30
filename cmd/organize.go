package cmd

import (
	"github.com/s0ders/go-torrent-organizer/internal/organizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(organizeCmd)
}

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize torrent downloaded files into Kodi folder structure",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return organizer.Organize(args[0], args[1])
	},
}
