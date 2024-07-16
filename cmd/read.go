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

var readInstanceId string

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read an instance by id",
	Long:  `read an instance by id - the long version!`,
	Run: func(cmd *cobra.Command, args []string) {
		gitopsClient, err := gitopsclient.NewGitopsClient(config)
		if err != nil {
			fmt.Println("GitopsClient error: ", err)
			os.Exit(1)
		}
		instance, err := gitopsClient.GetInstance(readInstanceId)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("INSTANCE: ", instance)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	readCmd.Flags().StringVar(&readInstanceId, "iid", "", "instance id as UUID string")
	readCmd.MarkFlagRequired("iid")
}
