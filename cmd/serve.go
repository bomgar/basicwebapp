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
		databaseUrl, _ := cmd.Flags().GetString("database-url")
		web.Run(web.RunSettings{
			ListenAddress: bindAddress,
			LogLevel:      logLevel,
			DatabaseUrl:   databaseUrl,
		})
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("bind-address", "i", ":8080", "Bind address")
	serveCmd.Flags().StringP("log-level", "l", "info", "Log level")
	serveCmd.Flags().StringP("database-url", "d", "postgres://fkbr:fkbr@localhost:5432/fkbr", "Bind address")

}
