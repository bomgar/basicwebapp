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
		port, _ := cmd.Flags().GetInt("port")
        web.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 8080, "Port to run the server on")

}
