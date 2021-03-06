/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package configuration

import (
	rootcmd "github.com/liangyuanpeng/nacosctl/cmd"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/spf13/cobra"
)

var (
	client      config_client.IConfigClient
	dataId      string
	group       string
	namespaceId string
)

// configsCmd represents the configs command
var configsCmd = &cobra.Command{
	Use:   "configs",
	Short: "This API is used to manager configurations in Nacos.",
	Long:  `This API is used to manager configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
		// fmt.Println("Hi.Here is Nacosctl 0.1.0 configsCmd")
	},
}

func init() {

	getConfigsCmd.PersistentFlags().StringVar(&dataId, "dataId", "", "dataId")
	getConfigsCmd.PersistentFlags().StringVar(&group, "group", "", "group")
	getConfigsCmd.PersistentFlags().StringVar(&namespaceId, "namespaceId", "public", "namespaceId")

	rootcmd.AddCommand(configsCmd)
	configsCmd.AddCommand(getConfigsCmd)
	configsCmd.AddCommand(createConfigsCmd)
}
