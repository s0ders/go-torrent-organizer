package cmd

import (
	"github.com/s0ders/go-torrent-organizer/internal/parser"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use:   "parser",
	Short: "Parses a torrent file name to extract informations",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return parser.Parse(args[0])
	},
}
