package configuration

import (
	"fmt"
	"log"

	rootcmd "github.com/liangyuanpeng/nacosctl/cmd"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
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

		var err error
		// 至少一个ServerConfig
		serverConfigs := []constant.ServerConfig{
			{
				IpAddr:      rootcmd.Host,
				ContextPath: "/nacos",
				Port:        rootcmd.Port,
				Scheme:      "http",
			},
		}

		//创建clientConfig
		clientConfig := constant.ClientConfig{
			NamespaceId:         namespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "tmp\\nacos\\log",
			CacheDir:            "tmp\\nacos\\cache",
			RotateTime:          "1h",
			MaxAge:              0,
			LogLevel:            "debug",
		}
		// 创建动态配置客户端
		client, err = clients.CreateConfigClient(map[string]interface{}{
			"serverConfigs": serverConfigs,
			"clientConfig":  clientConfig,
		})
		if err != nil {
			panic(err)
		}

		create(dataId, group, content, namespaceId)
	},
}

//get config from nacos
func create(dataId, group, content, namespaceId string) {
	success, err := client.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content})
	if err != nil {
		panic(err)
	}
	log.Println("success:", success)
}
