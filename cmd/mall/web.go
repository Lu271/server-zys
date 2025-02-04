package mall

import (
	"fmt"
	"github.com/Lu271/server-zys/api/router"
	"github.com/Lu271/server-zys/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
)

var config string
var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

func init() {
	rootCmd.AddCommand(webCmd)
	// 通过命令行参数传递配置文件路径
	webCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}

func startWebServer(cmd *cobra.Command, args []string) {
	fmt.Println(fmt.Sprintf("listen %v, start web server ......", core.GlobalConfig.Server.Addr))

	engine := gin.New()
	router.RegisterRouter(engine)
	server := initServer(engine)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func initServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         core.GlobalConfig.Server.Addr,
		Handler:      handler,
		ReadTimeout:  core.GlobalConfig.Server.ReadTimeout,
		WriteTimeout: core.GlobalConfig.Server.WriteTimeout,
		IdleTimeout:  core.GlobalConfig.Server.IdleTimeout,
	}
}
