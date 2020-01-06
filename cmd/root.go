package cmd

import (
	"fmt"
	"jarvis/internal/pkg/config"

	"github.com/spf13/cobra"
)

var (
	configFile string
	envName    string

	rootCmd = &cobra.Command{
		Use:   "jarvis",
		Short: "jarvis is our automation CLI",
		Long: `jarvis is the main command, used to facilitate SRE's life.
		
jarvis is smart, jarvis is beautiful,
jarvis is made with love by the SRE team.`,
		Run: func(cmd *cobra.Command, args []string) {
			//Here add random fun quotes from jarvis
			fmt.Println("Sir, take a deep breath.")
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVar(&configFile, "config", config.DEFAULT_CONFIG_PATH, "config file path")
	rootCmd.PersistentFlags().
		StringVarP(&envName, "env", "e", "", "Environment name")
}

func initConfig() {
	configReader := config.InitConfigurationReader(configFile)

	_, err := configReader.ParseConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error while parsing config file: %s \n", err))
	}

}

func Execute() error {
	return rootCmd.Execute()
}
