package cmd

import (
	"fmt"
	"net/http"
	"os"

	. "github.com/petewall/firmware-service/v2/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "firmware-service",
	Short: "A service for managing firmware binaries",
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		api := &API{
			FirmwareStore: &InMemoryFirmwareStore{},
			LogOutput:     cmd.OutOrStdout(),
		}

		port := viper.GetInt("port")
		cmd.Printf("Listening on port %d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), api.GetMux())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int("port", 5050, "Port to listen on")
	_ = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	_ = viper.BindEnv("port", "PORT")
}
