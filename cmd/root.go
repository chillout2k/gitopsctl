/*
Copyright © 2024 Dominik Chilla <dominik@zwackl.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/chillout2k/gitopsclient"
	"github.com/spf13/cobra"
)

var config gitopsclient.GitopsClientConfig
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitopsctl",
	Short: "gitops-client written in Go and Cobra",
	Long:  `gitops-client written in Go and Cobra - the long version!`,
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
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Trouble getting user homedir: ", err)
		os.Exit(1)
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&config.GitopsApiURI, "gau", "http://localhost:8000", "Gitops REST-API URI")
	rootCmd.PersistentFlags().StringVar(&config.JwksURI, "ju", "https://iam.internetservices.dev.datev.de/realms/chd/protocol/openid-connect/certs", "IDP JWKS URI")
	rootCmd.PersistentFlags().StringVar(&config.CachePath, "cp", homeDir+"/.gitopsctl", "path to persistent cache")
	rootCmd.PersistentFlags().BoolVar(&debug, "d", false, "enable verbose output")
}
