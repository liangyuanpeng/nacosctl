package configuration

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cobra"
)

var getConfigsCmd = &cobra.Command{
	Use:   "get",
	Short: "This API is used to get configurations in Nacos.",
	Long:  `This API is used to get configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 getConfigsCmd")

		if dataId == "" || group == "" {
			log.Println("dataId and group must be not null")
			cmd.Help()
			return
		}

		content, err := Get(dataId, group, namespaceId)
		if err != nil {
			log.Println("get config failed:", err)
			return
		}
		log.Println("content:", content)
	},
}

//get config from nacos
func Get(dataID, group, namespaceID string) (string, error) {
	//TODO Multiple namespace
	if client == nil {
		initClient()
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group})
	if err != nil {
		log.Println("get config.failed:", err)
		return content, err
	}

	return content, nil
}
