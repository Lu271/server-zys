package core

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
)

func InitSentinel() {
	// 初始化 Sentinel
	err := api.InitDefault()
	if err != nil {
		panic(err)
	}

	// 定义熔断规则
	rules := []*circuitbreaker.Rule{
		{
			Resource:         "example-resource",
			Strategy:         circuitbreaker.SlowRequestRatio,
			Threshold:        0.5,
			RetryTimeoutMs:   3000,
			MinRequestAmount: 10,
			StatIntervalMs:   60000,
			MaxAllowedRtMs:   500,
		},
	}

	// 加载熔断规则
	_, err = circuitbreaker.LoadRules(rules)
	if err != nil {
		panic(err)
	}
}
