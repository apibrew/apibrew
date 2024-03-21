package apbr

import (
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/spf13/cobra"
	"net/http"
)

var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"config"},
	Short:   "Configure apibrew",
	Long:    `Configure apibrew`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Configure apibrew\n")

		if err := client.AssureConfigFileExists(); err != nil {
			return err
		}

		client.LoadConfig()

		var config = client.GetConfig()

		config.Type = "ServerConfig"

		if config.DefaultServer == "" {
			config.DefaultServer = "default"
		}

		var serverConfig = client.ServerConfig{
			Name:           "default",
			Authentication: &client.ConfigServerAuthentication{},
		}

		cloud, err := cmd.Flags().GetBool("cloud")

		if err != nil {
			return err
		}

		local, err := cmd.Flags().GetBool("local")

		if err != nil {
			return err
		}

		if !cloud && !local {
			return errors.New("either cloud or local configuration must be chosen")
		}

		if cloud {
			projectName, err := cmd.Flags().GetString("project")

			if err != nil {
				return err
			}

			if projectName == "" {
				return errors.New("project name must be provided")
			}

			serverConfig.Host = projectName + ".apibrew.io"
			serverConfig.Port = 9443
			serverConfig.HttpPort = 8443

			res, err := http.Get("https://" + serverConfig.Host + ":" + fmt.Sprint(serverConfig.HttpPort) + "/health")

			if err != nil {
				return err
			}

			if res.StatusCode != 200 {
				return errors.New("project with name " + projectName + " does not exist or is not accessible")
			}
		} else {
			host, err := cmd.Flags().GetString("host")

			if err != nil {
				return err
			}

			if host == "" {
				return errors.New("host must be provided")
			}

			serverConfig.Host = host

			port, err := cmd.Flags().GetInt32("port")

			if err != nil {
				return err
			}

			if port == -1 {
				return errors.New("port must be provided")
			}

			serverConfig.Port = uint32(port)

			httpPort, err := cmd.Flags().GetInt32("httpPort")

			if err != nil {
				return err
			}

			if httpPort != -1 {
				serverConfig.HttpPort = uint32(httpPort)
			}

			insecure, err := cmd.Flags().GetBool("insecure")

			if err != nil {
				return err
			}

			serverConfig.Insecure = insecure
		}

		// authentication

		username, err := cmd.Flags().GetString("username")

		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")

		if err != nil {
			return err
		}

		token, err := cmd.Flags().GetString("token")

		if err != nil {
			return err
		}

		if username != "" && password != "" {
			serverConfig.Authentication.Username = username
			serverConfig.Authentication.Password = password
		} else if token != "" {
			serverConfig.Authentication.Token = token
		} else {
			return errors.New("either username and password or token must be provided")
		}

		serverName, err := cmd.Flags().GetString("server-name")

		if err != nil {
			return err
		}

		if serverName != "" {
			serverConfig.Name = serverName
		}

		config.Servers = append(config.Servers, serverConfig)

		err = client.WriteConfig()

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	configureCmd.PersistentFlags().Bool("cloud", false, "Cloud instance")
	configureCmd.PersistentFlags().String("project", "", "Cloud Project name")
	configureCmd.PersistentFlags().Bool("local", false, "Local instance")
	configureCmd.PersistentFlags().String("host", "", "ApiBrew instance Host address: e.g. localhost")
	configureCmd.PersistentFlags().Int32("port", -1, "ApiBrew instance Host port: e.g. 9009")
	configureCmd.PersistentFlags().Int32("httpPort", -1, "ApiBrew instance Host http port: e.g. 9009, http port is used by some services")
	configureCmd.PersistentFlags().Bool("insecure", true, "Insecure connection(ssl=false)")
	configureCmd.PersistentFlags().String("username", "", "Authentication / Username")
	configureCmd.PersistentFlags().String("password", "", "Authentication / Password")
	configureCmd.PersistentFlags().String("token", "", "Authentication / Token")
	configureCmd.PersistentFlags().String("server-name", "", "Server Name")
}
