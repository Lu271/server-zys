package mall

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"server-zys/internal/core"
	"server-zys/logs"
)

var rootCmd = &cobra.Command{
	Use:   "mall",
	Short: "CLI",
	Long:  "CLI for interacting with mall",
}

func init() {
	core.InitConfig(config)
	logs.InitLogger()
	core.InitSentinel()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
