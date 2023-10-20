package apbr

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
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
			Authentication: &client.ConfigServerAuthentication{},
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Do you want to connect to Api Brew cloud or your own server? (0 => Api Brew Cloud; 1 => Your own server): ")

		serverType, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		serverType = serverType[:len(serverType)-1]

		if serverType == "0" || serverType == "" {
			if err := configureCloudServerConfig(reader, &serverConfig); err != nil {
				return err
			}
		} else if serverType == "1" {
			if err := configureOwnServerConfig(reader, &serverConfig); err != nil {
				return err
			}
		} else {
			return errors.New("invalid server type")
		}

		for {
			fmt.Print("Enter Name[default]: ")
			serverName, err := reader.ReadString('\n')

			if err != nil {
				return err
			}

			serverName = serverName[:len(serverName)-1]

			if err != nil {
				return err
			}

			if serverName == "" {
				serverName = "default"
			}

			var serverNameExists = false
			for _, server := range config.Servers {
				if server.Name == serverName {
					fmt.Println("Server with name " + serverName + " already exists")
					serverNameExists = true
					break
				}
			}

			if !serverNameExists {
				serverConfig.Name = serverName
				break
			} else {
				fmt.Print("Do you want to overwrite the existing server? (0 => No; 1 => Yes) [0]: ")

				overwriteServerStr, err := reader.ReadString('\n')

				if err != nil {
					return err
				}

				overwriteServerStr = overwriteServerStr[:len(overwriteServerStr)-1]

				if overwriteServerStr == "0" || overwriteServerStr == "" {
					continue
				} else if overwriteServerStr == "1" {
					break
				} else {
					return errors.New("invalid option")
				}
			}
		}

		fmt.Println("Configuring authentication")

		fmt.Print("Which authentication method do you want to use? (0 => credentials; 1 => token) [0]: ")

		authMethodStr, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		authMethodStr = authMethodStr[:len(authMethodStr)-1]

		if authMethodStr == "0" || authMethodStr == "" {
			fmt.Print("Enter Username: ")
			_, err = fmt.Scanln(&serverConfig.Authentication.Username)

			if err != nil {
				return err
			}

			fmt.Print("Enter Password: ")
			_, err = fmt.Scanln(&serverConfig.Authentication.Password)

			if err != nil {
				return err
			}
		} else {
			fmt.Print("Enter Token: ")
			_, err = fmt.Scanln(&serverConfig.Authentication.Token)

			if err != nil {
				return err
			}
		}

		var servers []client.ServerConfig

		for _, server := range config.Servers {
			if server.Name == serverConfig.Name {
				continue
			}
			servers = append(servers, server)
		}

		servers = append(servers, serverConfig)

		config.Servers = servers

		err = client.WriteConfig()

		if err != nil {
			return err
		}

		fmt.Print("All done!\n")

		return nil
	},
}

func configureCloudServerConfig(reader *bufio.Reader, serverConfig *client.ServerConfig) error {
	for {
		fmt.Print("Enter Project ID (e.g. project-77fbc4): ")
		projectName, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		projectName = projectName[:len(projectName)-1]

		if projectName == "" {
			fmt.Println("Error: project name cannot be empty")
			continue
		}

		serverConfig.Host = projectName + ".apibrew.io"
		serverConfig.Port = 9443
		serverConfig.HttpPort = 8443

		res, err := http.Get("https://" + serverConfig.Host + ":" + fmt.Sprint(serverConfig.HttpPort) + "/health")

		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			fmt.Println("Error: Project with name " + projectName + " does not exist or is not accessible")
			continue
		}

		break
	}

	return nil
}

func configureOwnServerConfig(reader *bufio.Reader, serverConfig *client.ServerConfig) error {
	fmt.Print("Enter Host[localhost]: ")
	serverHost, err := reader.ReadString('\n')

	serverConfig.Host = serverHost[:len(serverHost)-1]

	if serverConfig.Host == "" {
		serverConfig.Host = "localhost"
	}

	if err != nil {
		return err
	}

	fmt.Print("Enter Port[9009]: ")

	serverPortStr, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	serverPortStr = serverPortStr[:len(serverPortStr)-1]

	if serverPortStr == "" {
		serverConfig.Port = 9009
	} else {
		port, err := strconv.Atoi(serverPortStr)

		if err != nil {
			return err
		}

		serverConfig.Port = uint32(port)
	}

	serverConfig.HttpPort = serverConfig.Port

	serverConfig.Insecure = true

	return nil
}
