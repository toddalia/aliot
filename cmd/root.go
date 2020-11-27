package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "remoteconsole",
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
	rootCmd.PersistentFlags().String("access-key", "", "阿里云帐号 AccessKey ID")
	rootCmd.PersistentFlags().String("access-secret", "", "阿里云帐号 AccessKey Secret")
	rootCmd.PersistentFlags().String("product-key", "", "物联网产品 ProduceKey")
	rootCmd.PersistentFlags().String("region", "cn-shanghai", "region ID")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 存储设置到viper
	viper.BindPFlag("access-key", rootCmd.PersistentFlags().Lookup("access-key"))
	viper.BindPFlag("access-secret", rootCmd.PersistentFlags().Lookup("access-secret"))
	viper.BindPFlag("product-key", rootCmd.PersistentFlags().Lookup("product-key"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.SetDefault("region", "cn-shanghai")
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
	if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // Config file not found; ignore error if desired
    } else {
			fmt.Printf("Fatal: %s\n", err)
			fmt.Printf("       config file: %s\n", viper.ConfigFileUsed())
			os.Exit(1)
    }
	}
}
