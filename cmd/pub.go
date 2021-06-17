package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toddalia/aliot/iot"
)

// 存放消息内容的文件路径
var msgFilePath string

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub deviceName",
	Short: "向设备发送消息",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		var deviceName = args[0]
		var message string

		device := &iot.Device{product, deviceName}

		if msgFilePath != "" {
			content, err := ioutil.ReadFile(msgFilePath)
			if err != nil {
				exitWithError(err)
			}
			message = string(content)
		}

		if viper.IsSet("message") {
			message = viper.GetString("message")
		}

		response, err := iot.Pub(client, device, message)
		if err != nil {
			exitWithError(err)
		}
		fmt.Print(response.GetHttpContentString())
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pubCmd.Flags().StringP("message", "m", "", "the message to send")
	pubCmd.Flags().StringVarP(&msgFilePath, "file", "f", "", "file that contains the message")

	viper.BindPFlag("message", pubCmd.Flags().Lookup("message"))
}
