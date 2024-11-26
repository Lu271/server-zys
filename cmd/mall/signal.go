package mall

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var signalCmd = &cobra.Command{
	Use: "signal",
	Run: startSignalServer,
}

func init() {
	rootCmd.AddCommand(signalCmd)
}

func startSignalServer(cmd *cobra.Command, args []string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-c
		fmt.Println("quit signal receive, quit")
		wg.Done()
	}()
	wg.Wait()
}
