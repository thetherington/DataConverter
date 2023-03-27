/*
Copyright © 2023 Tom Hetherington <thomas@hetheringtons.org>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thetherington/DataConverter/internal/app"
	"github.com/thetherington/DataConverter/internal/helpers"
	"github.com/thetherington/DataConverter/internal/log_ping"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping [flags] [archive]",
	Short: "This sub command converts inSITE ping index export from inSITE 10.3 to inSITE 11",
	Long: `Runing this sub command will extract the inSITE ping metric index export from inSITE version 10.3 and perform a conversion of the files so it's compatible with inSITE version 11
  
  Example: ./DataConverter ping log-metric-poller-ping-2022.03.31.tar.gz
	`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		index, err := helpers.ValidateArchiveGetIndexName(args[0])
		if err != nil {
			return err
		}

		converter := log_ping.New(index)

		app := app.New(converter)

		err = app.Run(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
