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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Gitops REST-API",
	Long: `Login to Gitops REST-API - get Oauth2 access-token from identity provider
	  and save it to ~/.gitopsctl/access_token`,
	Run: func(cmd *cobra.Command, args []string) {
		gitopsClient, err := gitopsclient.NewGitopsClient(config)
		if err != nil {
			fmt.Println("GitopsClient error: ", err)
			os.Exit(1)
		}
		err = gitopsClient.GetToken()
		if err != nil {
			fmt.Println("Login failed: ", err)
			os.Exit(1)
		}
		token, err := gitopsClient.ParseToken(gitopsClient.AccessToken)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		date, err := token.Claims.GetExpirationTime()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Login succeeded! Token expires on:", date.String())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVar(&config.ClientId, "cid", "gitops-playground", "Oauth2 client id")
	loginCmd.Flags().StringVar(&config.ClientSecret, "cs", "", "Oauth2 client secret (confidential client)")
	loginCmd.Flags().StringVar(&config.GrantType, "gt", "", "Oauth2 grant type (rfc6749): password, auth_code, device_code")
	loginCmd.MarkFlagRequired("gt")
	loginCmd.Flags().StringVar(&config.Scopes, "s", "openid", "Oauth2 scopes")
	loginCmd.Flags().StringVar(&config.TokenURI, "tu", "https://iam.internetservices.dev.datev.de/realms/chd/protocol/openid-connect/token", "Oauth2 token URI")
	loginCmd.Flags().StringVar(&config.Username, "user", "", "Oauth2 username (grant_type: password)")
	loginCmd.Flags().StringVar(&config.Password, "pass", "", "Oauth2 password (grant_type: password)")
	loginCmd.Flags().StringVar(&config.AuthURI, "au", "https://iam.internetservices.dev.datev.de/realms/chd/protocol/openid-connect/auth", "Oauth2 authorization URI (grant_type: authorization_code)")
	loginCmd.Flags().StringVar(&config.RedirectURI, "ru", "http://localhost:12345/authz", "Oauth2 redirect URI (grant_type: authorization_code)")
	loginCmd.Flags().StringVar(&config.AuthzListenerSocket, "als", "localhost:12345", "Oauth2 listener socket for local webserver (grant_type: authorization_code)")
}
