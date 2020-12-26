package configuration

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createConfigsCmd = &cobra.Command{
	Use:   "create",
	Short: "This API is used to create configurations in Nacos.",
	Long:  `This API is used to create configurations in Nacos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.Here is Nacosctl 0.1.0 createConfigsCmd")
	},
}
