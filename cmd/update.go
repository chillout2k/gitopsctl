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

var updateInstanceId string
var instanceUpdate gitopsclient.InstanceUpdate

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an instance by id",
	Long:  `update an instance by id - the long version!`,
	Run: func(cmd *cobra.Command, args []string) {
		gitopsClient, err := gitopsclient.NewGitopsClient(config)
		if err != nil {
			fmt.Println("GitopsClient error: ", err)
			os.Exit(1)
		}
		instance, err := gitopsClient.PutInstance(updateInstanceId, instanceUpdate)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("INSTANCE: ", instance)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVar(&updateInstanceId, "iid", "", "instance id")
	updateCmd.MarkFlagRequired("iid")
	updateCmd.Flags().StringVar(&instanceUpdate.Instance_name, "in", "", "instance name")
	updateCmd.MarkFlagRequired("in")
	updateCmd.Flags().Uint64Var(&instanceUpdate.Bits_account, "ba", 0, "BITS account")
	updateCmd.MarkFlagRequired("ba")
	updateCmd.Flags().Uint64Var(&instanceUpdate.Service_id, "sid", 0, "IT-service id")
	updateCmd.MarkFlagRequired("sid")
	updateCmd.Flags().Uint64Var(&instanceUpdate.Replica_count, "rc", 0, "replica count")
	updateCmd.MarkFlagRequired("rc")
	updateCmd.Flags().StringVar(&instanceUpdate.Version, "v", "", "product x version as semver constraint")
	updateCmd.MarkFlagRequired("v")
	updateCmd.Flags().StringVar(&instanceUpdate.Some_value, "sv", "", "some value")
	updateCmd.MarkFlagRequired("sv")
}
