package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/cobra"
	"youshupai.com/aliot/iot"
)

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "获取在线设备数量",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := sdk.NewClientWithAccessKey(product.Region, aliAccount.AccessKey, aliAccount.AccessSecret)
		if err != nil {
			exitWithError(err)
		}

		response, err := iot.QueryDeviceStatistics(client, product)
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
	rootCmd.AddCommand(statCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
