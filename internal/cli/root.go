package cli

import (
	"time"

	"github.com/KillAllChickens/gust/internal/weather"
	"github.com/spf13/cobra"
)

var update bool

var RootCmd = &cobra.Command{
	Use:     "config",
	Short:   "Set Config",
	Example: "urly config",
	Run: func(cmd *cobra.Command, args []string) {
		if update {
			for {
				weather.ShowWeather()
				time.Sleep(30 * time.Second)
			}
		}
		weather.ShowWeather()
	},
}

func init() {
	RootCmd.Flags().BoolVarP(&update, "update", "u", false, "Auto Update screen")
}
