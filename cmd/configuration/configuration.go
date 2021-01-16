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
	dataId      string
	group       string
	namespaceId string
	content     string
	client      config_client.IConfigClient
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

func initClient() {
	client = rootcmd.GetClient(namespaceId)
}

func init() {

	getConfigsCmd.PersistentFlags().StringVar(&dataId, "dataId", "", "dataId")
	getConfigsCmd.PersistentFlags().StringVar(&group, "group", "", "group")
	getConfigsCmd.PersistentFlags().StringVar(&namespaceId, "namespaceId", "", "namespaceId")

	deleteConfigsCmd.PersistentFlags().StringVar(&dataId, "dataId", "", "dataId")
	deleteConfigsCmd.PersistentFlags().StringVar(&group, "group", "", "group")
	deleteConfigsCmd.PersistentFlags().StringVar(&namespaceId, "namespaceId", "", "namespaceId")

	listenConfigsCmd.PersistentFlags().StringVar(&dataId, "dataId", "", "dataId")
	listenConfigsCmd.PersistentFlags().StringVar(&group, "group", "", "group")
	listenConfigsCmd.PersistentFlags().StringVar(&namespaceId, "namespaceId", "", "namespaceId")

	createConfigsCmd.PersistentFlags().StringVar(&dataId, "dataId", "", "dataId")
	createConfigsCmd.PersistentFlags().StringVar(&group, "group", "", "group")
	createConfigsCmd.PersistentFlags().StringVar(&namespaceId, "namespaceId", "", "namespaceId")
	createConfigsCmd.PersistentFlags().StringVar(&content, "content", "", "content")

	rootcmd.AddCommand(configsCmd)
	configsCmd.AddCommand(getConfigsCmd)
	configsCmd.AddCommand(createConfigsCmd)
	configsCmd.AddCommand(deleteConfigsCmd)
	configsCmd.AddCommand(listenConfigsCmd)
}
