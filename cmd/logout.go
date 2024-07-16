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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Gitops REST-API",
	Long:  `Logout from Gitops REST-API - Delete access/refresh token form file cache`,
	Run: func(cmd *cobra.Command, args []string) {
		gitopsClient, err := gitopsclient.NewGitopsClient(config)
		if err != nil {
			fmt.Println("GitopsClient error: ", err)
			os.Exit(1)
		}
		err = gitopsClient.DeleteTokenFromCache("access")
		if err != nil {
			fmt.Println("Logout failed: ", err)
			os.Exit(1)
		}
		err = gitopsClient.DeleteTokenFromCache("refresh")
		if err != nil {
			fmt.Println("Logout failed: ", err)
			os.Exit(1)
		}
		fmt.Println("Logout succeeded!")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
