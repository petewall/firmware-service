package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	. "github.com/petewall/firmware-service/v2/internal"
	. "github.com/petewall/firmware-service/v2/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	StoreTypeInMemory   = "memory"
	StoreTypeFilesystem = "filesystem"
)

func ValidateFirmwareStoreArgs(cmd *cobra.Command, args []string) error {
	firmwareStoreType := viper.GetString("store.type")
	if firmwareStoreType == "" {
		return errors.New("firmware store type is required. Please run again with --firmware-store-type")
	}

	if firmwareStoreType == StoreTypeInMemory {
		return nil
	}

	if firmwareStoreType == StoreTypeFilesystem {
		firmwareStorePath := viper.GetString("store.path")
		if firmwareStorePath == "" {
			return fmt.Errorf("filesystem firmware store requires a path. Please run again with --firmware-store-path")
		}

		// TODO: validate path is a valid directory

		return nil
	}

	return fmt.Errorf("unknown firmware store type: %s. Valid options are \"%s\" or \"%s\"", firmwareStoreType, StoreTypeInMemory, StoreTypeFilesystem)
}

var rootCmd = &cobra.Command{
	Use:     "firmware-service",
	Short:   "A service for managing firmware binaries",
	PreRunE: ValidateFirmwareStoreArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var firmwareStore FirmwareStore
		if viper.GetString("store.type") == StoreTypeInMemory {
			firmwareStore = NewInMemoryFirmwareStore()
		} else if viper.GetString("store.type") == StoreTypeFilesystem {
			firmwareStore = &FilesystemFirmwareStore{
				Path: viper.GetString("store.path"),
			}
		}

		api := &API{
			FirmwareStore: firmwareStore,
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

	rootCmd.Flags().String("firmware-store-type", "memory", "Type of firmware store to use.")
	_ = viper.BindPFlag("store.type", rootCmd.Flags().Lookup("firmware-store-type"))
	_ = viper.BindEnv("store.type", "FIRMWARE_STORE_TYPE")

	rootCmd.Flags().String("firmware-store-path", "", "Path for file system firmware store.")
	_ = viper.BindPFlag("store.path", rootCmd.Flags().Lookup("firmware-store-path"))
	_ = viper.BindEnv("store.path", "FILESYSTEM_FIRMWARE_STORE_PATH")

	rootCmd.SetOut(rootCmd.OutOrStdout())
}
