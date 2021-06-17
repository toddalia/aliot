package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toddalia/aliot/iot"
)

// 配置文件路径
var cfgFile string
var product *iot.Product
var client *sdk.Client

var rootCmd = &cobra.Command{
	Use:   "aliot",
	Short: "树莓派远程调试启动工具",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is config.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "config.json"
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName("config")
	}

	// If a config file is found, read it in.
	err := viper.ReadInConfig();
	if err != nil {
		fmt.Printf("Fatal: %s\n", err)
		fmt.Printf("       config file: %s\n", viper.ConfigFileUsed())
		os.Exit(1)
	}

	if !(viper.IsSet("accessKey") && viper.IsSet("accessSecret")) {
		exitWithError(errors.New("请在配置文件里设置阿里云账户: accessKey, accessSecret"))
	}

	aliAccount := &iot.AliAccount{
		AccessKey: viper.GetString("accessKey"),
		AccessSecret: viper.GetString("accessSecret"),
	}

	if !viper.IsSet("productKey") {
		exitWithError(errors.New("请在配置文件里设置 productKey"))
	}

	if !viper.IsSet("region") {
		exitWithError(errors.New("请在配置文件里设置 region"))
	}

	product = &iot.Product{
		ProductKey: viper.GetString("productKey"),
		Region: viper.GetString("region"),
	}

	client, err = sdk.NewClientWithAccessKey(product.Region, aliAccount.AccessKey, aliAccount.AccessSecret)
	if err != nil {
		exitWithError(err)
	}
}
