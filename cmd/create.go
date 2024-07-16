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

var instanceOrder = gitopsclient.InstanceOrder{}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new instance",
	Long:  `create a new instance - the long version!`,
	Run: func(cmd *cobra.Command, args []string) {
		gitopsClient, err := gitopsclient.NewGitopsClient(config)
		if err != nil {
			fmt.Println("GitopsClient error: ", err)
			os.Exit(1)
		}
		instance, err := gitopsClient.PostInstanceOrder(instanceOrder)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("New instance: ", instance.Instance_id)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&instanceOrder.Instance_name, "in", "", "instance name")
	createCmd.MarkFlagRequired("in")
	createCmd.Flags().StringVar(&instanceOrder.Orderer_id, "oid", "", "orderer id")
	createCmd.MarkFlagRequired("oid")
	createCmd.Flags().Uint64Var(&instanceOrder.Bits_account, "ba", 0, "BITS account")
	createCmd.MarkFlagRequired("ba")
	createCmd.Flags().Uint64Var(&instanceOrder.Service_id, "sid", 0, "IT-service id")
	createCmd.MarkFlagRequired("sid")
	createCmd.Flags().Uint64Var(&instanceOrder.Replica_count, "rc", 0, "replica count")
	createCmd.MarkFlagRequired("rc")
	createCmd.Flags().StringVar(&instanceOrder.Version, "v", "", "product x version as semver constraint")
	createCmd.MarkFlagRequired("v")
	createCmd.Flags().StringVar(&instanceOrder.Some_value, "sv", "", "some value")
	createCmd.MarkFlagRequired("sv")
}
