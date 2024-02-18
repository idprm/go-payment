package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * Load config
		 */
		// cfg, err := config.LoadSecret("secret.yaml")
		// if err != nil {
		// 	panic(err)
		// }

		// lg := logger.NewLogger(cfg)

		// p := ximpay.NewXimpay(cfg, lg, &entity.Application{}, &entity.Gateway{}, &entity.Channel{}, &entity.Order{}, &entity.Payment{}, ve)
		// p.Payment()
	},
}
