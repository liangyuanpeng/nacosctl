package configuration

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cobra"
)

var createConfigsCmd = &cobra.Command{
	Use:   "publish",
	Short: "This API is used to create configurations in Nacos.",
	Long:  `This API is used to create configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 createConfigsCmd")

		if dataId == "" || group == "" || content == "" {
			log.Println("dataId and group and content must be not null")
			cmd.Help()
			return
		}

		initClient()

		err := create(dataId, group, content, namespaceId)
		if err != nil {
			log.Println("create config failed!", err)
		}
	},
}

//get config from nacos
func create(dataID, group, content, namespaceID string) error {
	success, err := client.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content})
	if err != nil {
		return err
	}
	log.Println("success:", success)
	return nil
}
