package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/KillAllChickens/gust/internal/config"
	"github.com/KillAllChickens/gust/internal/geocode"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "Set Config",
	Example: "urly config",
	Run: func(cmd *cobra.Command, args []string) {
		savedConfig := config.Default()

		unitsAsk, _ := Ask("Which units? (M)etric/(I)mperial: ")
		units := strings.ToLower(strings.TrimSpace(unitsAsk))

		switch units {
		case "i", "imperial":
			savedConfig.Units = "imperial"
		case "m", "metric":
			savedConfig.Units = "metric"
		default:
			savedConfig.Units = "metric"
		}

		city, _ := Ask("Find your city: ")
		cities := geocode.SearchCity(city)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tSTATE/REGION\tCOUNTRY\tTIMEZONE")
		for i, res := range cities.Results {
			fmt.Fprintf(w, "[%d]\t%s\t%s\t%s\t%s\n", i+1, res.Name, res.State, res.CountryCode, res.Timezone)
		}
		w.Flush()
		cityIDStr, _ := Ask("What ID is associated with your city? ")

		cityID, _ := strconv.Atoi(cityIDStr)

		savedConfig.Lat = cities.Results[cityID-1].Lat
		savedConfig.Lon = cities.Results[cityID-1].Long
		savedConfig.Location = cities.Results[cityID-1].Name

		savedConfig.Save()
	},
}
