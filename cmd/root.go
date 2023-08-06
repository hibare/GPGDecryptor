package cmd

import (
	"fmt"
	"os"

	"github.com/hibare/GPGDecryptor/internal/gpg"
	"github.com/hibare/GPGDecryptor/internal/version"
	"github.com/spf13/cobra"
)

var (
	privateKeyPath string
	publicKeyPath  string
	inputGPGFile   string

	rootCmd = &cobra.Command{
		Use:     "GPGDecryptor",
		Version: version.CurrentVersion,
		Short:   "Program to decrypt GPG files",
		Long:    ``,
		Run: func(cmd *cobra.Command, args []string) {
			gpg.GPGDecryptor(privateKeyPath, publicKeyPath, inputGPGFile)
			v := version.GetNewVersionInfo()
			if v.NewVersionAvailable && v.LatestVersion != "" {
				fmt.Printf("\n%s\n", v.GetUpdateNotification())
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&privateKeyPath, "priv-key", "", "Private key file path")
	rootCmd.PersistentFlags().StringVar(&publicKeyPath, "pub-key", "", "Public key file path")
	rootCmd.PersistentFlags().StringVar(&inputGPGFile, "gpg-file", "", "Input GPG file")
	rootCmd.MarkPersistentFlagRequired("priv-key")
	rootCmd.MarkPersistentFlagRequired("pub-key")
	rootCmd.MarkPersistentFlagRequired("gpg-file")
}
