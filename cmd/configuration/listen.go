package configuration

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cobra"
)

var listenConfigsCmd = &cobra.Command{
	Use:   "listen",
	Short: "This API is used to listen configurations in Nacos.",
	Long:  `This API is used to listen configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 listenConfigsCmd")

		if dataId == "" || group == "" {
			log.Println("dataId and group must be not null")
			cmd.Help()
			return
		}

		initClient()

		listen(dataId, group, namespaceId)
	},
}

//listen config from nacos
func listen(dataID, group, namespaceID string) {
	//TODO need sdk support custom listen time
	//TODO support canel listen
	err := client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			log.Printf("ns:%s group:%s dataId:%s data:%s", namespace, group, dataId, data)
		},
	})
	if err != nil {
		panic(err)
	}
	// time.Sleep(time.Duration(10) * time.Second)
}
