package command

import (
	"os"

	"github.com/owenthereal/upterm/server"
	"github.com/owenthereal/upterm/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flagSSHAddr              string
	flagWSAddr               string
	flagNodeAddr             string
	flagAllowCustomSessionID bool
	flagPrivateKeys          []string
	flagHostnames            []string
	flagNetwork              string
	flagNetworkOpts          []string
	flagMetricAddr           string
	flagDebug                bool
)

func Root(logger log.FieldLogger) *cobra.Command {
	rootCmd := &rootCmd{}
	cmd := &cobra.Command{
		Use:   "uptermd",
		Short: "Upterm Daemon",
		RunE:  rootCmd.Run,
	}

	cmd.PersistentFlags().StringVarP(&flagSSHAddr, "ssh-addr", "", utils.DefaultLocalhost("2222"), "ssh server address")
	cmd.PersistentFlags().StringVarP(&flagWSAddr, "ws-addr", "", "", "websocket server address")
	cmd.PersistentFlags().StringVarP(&flagNodeAddr, "node-addr", "", "", "node address")
	cmd.PersistentFlags().BoolVarP(&flagAllowCustomSessionID, "allow-custom-session-id", "", false, "if true, respect the custom session ID given by the hosts")
	cmd.PersistentFlags().StringSliceVarP(&flagPrivateKeys, "private-key", "", nil, "server private key")
	cmd.PersistentFlags().StringSliceVarP(&flagHostnames, "hostname", "", nil, "server hostname for public-key authentication certificate principals. If empty, public-key authentication is used instead.")

	cmd.PersistentFlags().StringVarP(&flagNetwork, "network", "", "mem", "network provider")
	cmd.PersistentFlags().StringSliceVarP(&flagNetworkOpts, "network-opt", "", nil, "network provider option")

	cmd.PersistentFlags().StringVarP(&flagMetricAddr, "metric-addr", "", "", "metric server address")
	cmd.PersistentFlags().BoolVarP(&flagDebug, "debug", "", os.Getenv("DEBUG") != "", "debug")

	return cmd
}

type rootCmd struct {
}

func (cmd *rootCmd) Run(c *cobra.Command, args []string) error {
	opt := server.Opt{
		SSHAddr:              flagSSHAddr,
		WSAddr:               flagWSAddr,
		NodeAddr:             flagNodeAddr,
		AllowCustomSessionID: flagAllowCustomSessionID,
		KeyFiles:             flagPrivateKeys,
		Hostnames:            flagHostnames,
		Network:              flagNetwork,
		NetworkOpt:           flagNetworkOpts,
		MetricAddr:           flagMetricAddr,
		Debug:                flagDebug,
	}

	return server.Start(opt)
}
