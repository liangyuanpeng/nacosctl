package configuration

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cobra"
)

var deleteConfigsCmd = &cobra.Command{
	Use:   "delete",
	Short: "This API is used to delete configurations in Nacos.",
	Long:  `This API is used to delete configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 deleteConfigsCmd")

		if dataId == "" || group == "" {
			log.Println("dataId and group must be not null")
			cmd.Help()
			return
		}
		initClient()

		delete(dataId, group, namespaceId)
	},
}

//delete config from nacos
func delete(dataID, group, namespaceID string) {
	success, err := client.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	//Tenant: namespaceId
	if err != nil {
		panic(err)
	}
	log.Println("success:", success)
}
