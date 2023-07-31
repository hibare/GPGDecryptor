/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/hibare/GPGDecryptor/pkg/gpg"
	"github.com/spf13/cobra"
)

var (
	privateKeyPath string
	publicKeyPath  string
	inputGPGFile   string

	rootCmd = &cobra.Command{
		Use:   "GPGDecryptor",
		Short: "Program to decrypt GPG files",
		Long:  ``,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			gpg.GPGDecryptor(privateKeyPath, publicKeyPath, inputGPGFile)
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
