package configuration

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var historyConfigsCmd = &cobra.Command{
	Use:   "history",
	Short: "This API is used to get history configurations in Nacos.",
	Long:  `This API is used to get history configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 get history Configs Cmd")

		if dataId == "" || group == "" {
			log.Println("dataId and group must be not null")
			cmd.Help()
			return
		}

		content, err := History(dataId, group, namespaceId)
		if err != nil {
			log.Println("get history config failed:", err)
			return
		}
		log.Println("content:", content)
	},
}

func History(dataID, group, namespaceID string) (string, error) {
	if client == nil {
		initClient()
	}
	//todo
	return content, nil
}

type ConfigHistory struct {
	PageItems []ConfigDto `json:"pageItems"`
}

type ConfigDto struct {
	dataId           string
	group            string
	lastModifiedTime string
}
