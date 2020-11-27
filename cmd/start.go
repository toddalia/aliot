package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"youshupai.com/remoteconsole/iot"
)

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [flags] deviceName",
	Short: "向树莓派发送指令，启动frp服务",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		var (
			deviceName = args[0]
			accessKey = viper.GetString("access-key")
			accessSecret = viper.GetString("access-secret")
			region = viper.GetString("region")
			productKey = viper.GetString("product-key")
			serverAddr = viper.GetString("server-addr")
			serverPort = viper.GetString("server-port")
			remotePort = viper.GetString("remote-port")
			token = viper.GetString("token")
		)

		if !(viper.IsSet("access-key") && viper.IsSet("access-secret")) {
			exitWithError(errors.New("未指定阿里云账户"))
		}

		device := &iot.Device{
			Product: iot.Product{
				ProductKey: productKey,
				Region: region,
			},
			Name: deviceName,
		}

		client, err := sdk.NewClientWithAccessKey(region, accessKey, accessSecret)
		if err != nil {
			exitWithError(err)
		}

		message := iot.NewMessage("remoteconsole", map[string]string {
			"server_addr": serverAddr,
			"server_port": serverPort,
			"remote_port": remotePort,
			"token": token,
		})

		response, err := iot.Pub(client, device, message)
		if err != nil {
			exitWithError(err)
		}
		fmt.Print(response.GetHttpContentString())
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  startCmd.Flags().StringP("server-addr", "a", "", "FRP 服务的地址")
	startCmd.Flags().StringP("server-port", "p", "", "FRP 服务的端口")
	startCmd.Flags().StringP("remote-port", "r", "", "FRP 服务的端口")
	startCmd.Flags().StringP("token", "t", "", "FRP 服务的token")

	// 保存设置到 viper
	viper.BindPFlag("server-addr", startCmd.Flags().Lookup("server-addr"))
	viper.BindPFlag("server-port", startCmd.Flags().Lookup("server-port"))
	viper.BindPFlag("remote-port", startCmd.Flags().Lookup("remote-port"))
	viper.BindPFlag("token", startCmd.Flags().Lookup("token"))
	viper.SetDefault("server-addr", "139.224.106.207")
}