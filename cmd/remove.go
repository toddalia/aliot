package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toddalia/aliot/iot"
)

// removeCmd 删除指定的设备
var removeCmd = &cobra.Command{
	Use:   "remove deviceName",
	Short: "删除设备",
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

		response, err := iot.DeleteDevice(client, device)
		if err != nil {
			exitWithError(err)
		}
		fmt.Print(response.GetHttpContentString())
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
