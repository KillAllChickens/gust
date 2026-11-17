package cli

import (
	"github.com/KillAllChickens/gust/internal/weather"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "config",
	Short:   "Set Config",
	Example: "urly config",
	Run: func(cmd *cobra.Command, args []string) {
		weather.ShowWeather()
	},
}
