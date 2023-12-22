package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func init() {
	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)

	/**
	 * WORKER SERVICE
	 */
	rootCmd.AddCommand(workerCmd)
	rootCmd.AddCommand(seederCmd)

	rootCmd.AddCommand(testCmd)

}

func Execute() error {
	return rootCmd.Execute()
}
