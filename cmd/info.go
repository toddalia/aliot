package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/cobra"
	"github.com/toddalia/aliot/iot"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info deviceName",
	Short: "获取设备信息",
	Run: func(cmd *cobra.Command, args []string) {
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

		response, err := iot.QueryDeviceDetail(client, device)
		if err != nil {
			exitWithError(err)
		}

		var respJSON iot.Response

		json.Unmarshal([]byte(response.GetHttpContentString()), &respJSON)

		for key, val := range respJSON.Data {
			fmt.Printf("%s\t%v\n", key, val)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
