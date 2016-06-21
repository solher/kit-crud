package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	vip = viper.New()
)

var (
	cfgFile string
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kit-crud",
	Short: "A CRUD go-kit based microservice.",
}

func init() {
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which if defined here will be global for your application

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (defaults are $HOME/config.yaml and ./config.yaml)")
	vip.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// Read in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		vip.SetConfigFile(cfgFile)
	}

	vip.SetConfigName("config") // name of config file (without extension)
	vip.AddConfigPath("$HOME")  // adding home directory as first search path
	vip.AddConfigPath("./")     // adding local directory as second search path
	vip.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := vip.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", vip.ConfigFileUsed())
	}
}

func Execute() error {
	return rootCmd.Execute()
}
