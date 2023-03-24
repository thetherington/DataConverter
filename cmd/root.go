/*
Copyright © 2023 Tom Hetherington <thomas@hetheringtons.org>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DataConverter",
	Short: "Convert inSITE v10.3 Export Data to inSITE 11 Data",
	Long:  `This application will convert the export data from elasticdump for data stored Elasticsearch 5.5 to Elasticsearch 7.16`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// err = doc.GenMarkdownTree(rootCmd, "docs")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.DataConverter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = "0.1"
}
