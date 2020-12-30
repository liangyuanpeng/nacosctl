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
package console

import (
	"bufio"
	"log"
	"os"
	"strings"

	rootcmd "github.com/liangyuanpeng/nacosctl/cmd"
	"github.com/liangyuanpeng/nacosctl/cmd/configuration"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/spf13/cobra"
)

var client config_client.IConfigClient

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "This API is console for Nacos.",
	Long:  `This API is console for Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("you are console ,input pls,\n 1) input exit for exit. \n 2) input help for base help info. \n 3) more interesting coming... ")
		in := bufio.NewScanner(os.Stdin)
		var currentPattern string
		for in.Scan() {
			inpu := in.Text()

			if inpu == "" {
				continue
			}

			if currentPattern != "" {
				if currentPattern == "get" {
					strs := strings.Split(inpu, "/")
					if len(strs) > 3 || len(strs) < 2 {
						log.Println("format have some wrong,print {namespace}/{group}/dataId to get config")
						continue
					}
					namespaceID := strs[0]
					var group string
					var dataID string
					if len(strs) == 2 {
						namespaceID = "public"
						group = strs[0]
						dataID = strs[1]
					} else {
						group = strs[1]
						dataID = strs[2]
					}

					// if client == nil {
					// 	client = rootcmd.GetClient(namespaceID)
					// }
					log.Printf("get config param:%s|%s|%s", namespaceID, group, dataID)
					content, err := configuration.Get(dataID, group, namespaceID)
					if err == nil {
						log.Printf("get config:%s", content)
					}
					continue
				}
			}

			switch inpu {
			case "exit":
				break
			case "help":
				cmd.Help()
				//support get publish delete listen
			case "get":
				currentPattern = "get"
				log.Printf("===choose get time,print {namespace}/{group}/dataId to get config,like public/lan.io/appv===")
				continue
			case "reset":
				currentPattern = ""
				log.Printf(">>>reset pattern")
				continue
			case "current":
				log.Printf(">>>print current pattern:%s", currentPattern)
				continue
			}
			log.Printf("get input:%s", inpu)
		}
	},
}

func init() {
	rootcmd.AddCommand(consoleCmd)
}
