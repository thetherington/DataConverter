/*
Copyright © 2023 Tom Hetherington <thomas@hetheringtons.org>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thetherington/DataConverter/internal/helpers"
	"github.com/thetherington/DataConverter/internal/log_syslog"
)

// syslogCmd represents the syslog command
var syslogCmd = &cobra.Command{
	Use:   "syslog",
	Short: "This sub command converts inSITE syslog index exports from inSITE 10.3 to inSITE 11",
	Long: `Runing this sub command will extract the inSITE syslog index export from inSITE version 10.3 and perform a conversion of the files so it's compatible with inSITE version 11
  
  Example: ./DataConverter syslog log-syslog-2022.03.31.tar.gz
	`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sm, spinners := helpers.CreateSpinGroups()

		// start the spinners
		sm.Start()

		index, err := helpers.ValidateArchiveGetIndexName(args[0])
		if err != nil {
			spinners["setup"].Error()
			sm.Stop()
			return err
		}

		err = helpers.CreateWorkingDirExtract(args[0], index)
		if err != nil {
			spinners["setup"].Error()
			sm.Stop()
			return err
		}

		spinners["setup"].UpdateMessage("Setup/Extraction...Complete")
		spinners["setup"].Complete()

		s := log_syslog.New(index)

		helpers.CreateTemplates(filepath.Join(index, "v11"), s)

		helpers.ScanAndConvert(
			filepath.Join(index, "v10.3", fmt.Sprintf("%s-data.json", index)),
			filepath.Join(index, "v11", fmt.Sprintf("%s-data.json", index)),
			s,
		)

		if s.ReturnErrors() == 0 {
			spinners["upgrade"].UpdateMessage("Upgrading Data...Complete")
		} else {
			spinners["upgrade"].UpdateMessagef("Upgrading Data...Errors: %d", s.ReturnErrors())
			spinners["upgrade"].Error()
		}

		spinners["upgrade"].Complete()

		err = helpers.ArchiveCleanup(index)
		if err != nil {
			spinners["cleanup"].Error()
			sm.Stop()
			return err
		}

		spinners["cleanup"].UpdateMessage("Archiving/Cleanup...Complete")
		spinners["cleanup"].Complete()

		sm.Stop()

		fmt.Println()
		fmt.Println("New Archive: ", filepath.Join(index, fmt.Sprintf("%s.tar.gz", index)))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(syslogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syslogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syslogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
