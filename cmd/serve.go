package cmd

import (
	"github.com/bomgar/basicwebapp/web"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server.`,
	Run: func(cmd *cobra.Command, _ []string) {
		bindAddress, _ := cmd.Flags().GetString("bind-address")
		logLevel, _ := cmd.Flags().GetString("log-level")
		web.Run(web.RunSettings{
			ListenAddress: bindAddress,
			LogLevel:      logLevel,
		})
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("bind-address", "i", ":8080", "Bind address")
	serveCmd.Flags().StringP("log-level", "l", "info", "Log level")

}
