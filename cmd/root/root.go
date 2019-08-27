package root

import (
	"github.com/kubemq-io/kubetools/cmd/logs"
	"github.com/kubemq-io/kubetools/cmd/proxy"
	"github.com/kubemq-io/kubetools/cmd/queue"
	"github.com/kubemq-io/kubetools/pkg/config"
	"github.com/kubemq-io/kubetools/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg *config.Config
var version string
var rootCmd = &cobra.Command{
	Use: "kubetools",
}

func Execute(ver string) {
	version = ver
	defer utils.CheckErr(cfg.Save())
	utils.CheckErr(rootCmd.Execute())

}

func init() {
	cfg = &config.Config{}
	defaultCfg, err := config.CheckConfigFile()
	if err != nil && defaultCfg != nil {
		cfg = defaultCfg
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(".kubetools")
		err := viper.ReadInConfig()
		utils.CheckErr(err)
		err = viper.Unmarshal(cfg)
		utils.CheckErr(err)
	}
	rootCmd.AddCommand(queue.NewCmdQueue(cfg))
	rootCmd.AddCommand(logs.NewCmdLogs(cfg))
	rootCmd.AddCommand(proxy.NewCmdProxy(cfg))
}