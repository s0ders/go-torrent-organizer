package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-torrent-organizer",
	Short: "A tool to organize your torrent movies and shows",
	Long: `go-torrent-organizer is a CLI meant to be used as a hook by
	your torrent client to automatically rename and organize your downloaded
	movies or tv shows.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".torrent-organizer")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
