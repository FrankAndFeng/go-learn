package rpcServer

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	StartCmd  = &cobra.Command{
		Use:     "rpcServer",
		Short:   "Initialize the database",
		Example: "go-admin migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	fmt.Println("init rpc server")
}

func run(){
	fmt.Println("run rpc server")
}