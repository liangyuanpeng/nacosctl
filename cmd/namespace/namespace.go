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
package namespace

import (
	rootcmd "github.com/liangyuanpeng/nacosctl/cmd"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/spf13/cobra"
)

var (
	client config_client.IConfigClient
)

// namespaceCmd represents the namespaces command
var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "This API is used to manager namespaces in Nacos.",
	Long:  `This API is used to manager namespaces in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootcmd.AddCommand(namespacesCmd)
}
