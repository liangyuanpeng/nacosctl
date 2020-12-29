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

		listen(dataId, group, namespaceId)
	},
}

//listen config from nacos
func listen(dataId, group, namespaceId string) {
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
