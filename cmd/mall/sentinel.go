package mall

import (
	"context"
	"fmt"
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/util"
	"github.com/spf13/cobra"
	"server-zys/logs"
	"time"
)

var sentinelCmd = &cobra.Command{
	Use: "sentinel",
	Run: startSentinelServer,
}

func init() {
	rootCmd.AddCommand(sentinelCmd)
}

func startSentinelServer(cmd *cobra.Command, args []string) {
	err := api.InitDefault()
	if err != nil {
		logs.Error(context.Background(),
			fmt.Sprintf("Sentinel initialization failed: %v", err))
		return
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			Threshold:              10,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		logs.Error(context.Background(),
			fmt.Sprintf("Failed to load rules: %v", err))
	}

	for i := 0; i < 10; i++ {
		go func(index int) {
			for {
				entry, blockError := api.Entry("some-test", api.WithTrafficType(base.Inbound))
				if blockError != nil {
					// 请求被限流
					fmt.Printf("Request #%d was blocked!\n", index)
				} else {
					// 请求成功通过
					fmt.Println(util.CurrentTimeMillis(), "Passed")
					entry.Exit() // 退出 Sentinel Entry
				}
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
}
