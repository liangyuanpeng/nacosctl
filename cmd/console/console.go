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

	rootcmd "github.com/liangyuanpeng/nacosctl/cmd"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "This API is console for Nacos.",
	Long:  `This API is console for Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		//使用os.Stdin开启输入流
		log.Println("you are console ,input pls,\n 1) input exit for exit. \n 2) input help for base help info. \n 3) more interesting coming... ")
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			inpu := in.Text()
			if inpu == "exit" {
				break
			}
			if inpu == "help" {
				cmd.Help()
			}
			log.Printf("get input:%s", inpu)
		}
	},
}

func init() {
	rootcmd.AddCommand(consoleCmd)
}
