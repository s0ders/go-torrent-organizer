package cmd

import (
	"fmt"

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
		infos, err := parser.ParseToJSON(args[0])
		if err != nil {
			return err
		}

		fmt.Println(infos)

		return nil
	},
}
