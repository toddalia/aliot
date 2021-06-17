package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toddalia/aliot/iot"
)

var frpCmd = &cobra.Command{
	Use:   "frp [flags] deviceName",
	Short: "向树莓派发送指令，启动frp服务",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		var (
			deviceName = args[0]
			serverAddr = viper.GetString("serverAddr")
			serverPort = viper.GetString("serverPort")
			remotePort = viper.GetString("remotePort")
			token = viper.GetString("token")
		)

		device := &iot.Device{
			Product: product,
			Name: deviceName,
		}

		message := iot.NewMessage("remoteconsole", map[string]string {
			"server_addr": serverAddr,
			"server_port": serverPort,
			"remote_port": remotePort,
			"token": token,
		})

		response, err := iot.PubMessage(client, device, message)
		if err != nil {
			exitWithError(err)
		}
		fmt.Print(response.GetHttpContentString())
	},
}

func init() {
	rootCmd.AddCommand(frpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  frpCmd.Flags().StringP("serverAddr", "a", "", "FRP 服务的地址")
	frpCmd.Flags().StringP("serverPort", "p", "", "FRP 服务的端口")
	frpCmd.Flags().StringP("remotePort", "r", "", "FRP 服务的端口")
	frpCmd.Flags().StringP("token", "t", "", "FRP 服务的token")

	// 保存设置到 viper
	viper.BindPFlag("serverAddr", frpCmd.Flags().Lookup("serverAddr"))
	viper.BindPFlag("serverPort", frpCmd.Flags().Lookup("serverPort"))
	viper.BindPFlag("remotePort", frpCmd.Flags().Lookup("remotePort"))
	viper.BindPFlag("token", frpCmd.Flags().Lookup("token"))
}
