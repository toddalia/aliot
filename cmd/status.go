package cmd

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/cobra"
	"github.com/toddalia/aliot/iot"
)

var statusCmd = &cobra.Command{
	Use: "status deviceName",
	Short: "查询设备状态",
	Run: func(cmd *cobra.Command, args []string)  {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		var (
			deviceName = args[0]
		)

		device := &iot.Device{
			Product: product,
			Name: deviceName,
		}

		client, err := sdk.NewClientWithAccessKey(product.Region, aliAccount.AccessKey, aliAccount.AccessSecret)
		if err != nil {
			exitWithError(err)
		}

		response, err := iot.GetDeviceStatus(client, device)
		if err != nil {
			exitWithError(err)
		}
		fmt.Print(response.GetHttpContentString())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
