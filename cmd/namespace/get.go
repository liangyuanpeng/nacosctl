package namespace

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getNamespaceCmd = &cobra.Command{
	Use:   "get",
	Short: "This API is used to get namespace in Nacos.",
	Long:  `This API is used to get namespace in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO need sdk support manager namespace
		fmt.Println("Hi.Here is Nacosctl 0.1.0 get namespaces Cmd")
	},
}
