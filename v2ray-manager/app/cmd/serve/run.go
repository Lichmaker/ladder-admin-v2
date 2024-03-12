package serve

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"

	serveController "v2ray-manager/app/controller/serve"
	"v2ray-manager/pkg/config"
	"v2ray-manager/pkg/logger"
	"v2ray-manager/protobuf/manager"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runServe,
	Args:  cobra.NoArgs,
}

func init() {
	// 注册子命令
	CmdServe.AddCommand(
	// CmdRestart,
	)
}

// 是否以daemon形式启动,默认为false
var RunServerDaemon bool

func buildContext() *daemon.Context {
	return &daemon.Context{
		PidFileName: "v2ray-manager.pid",
		PidFilePerm: 0644,
		LogFileName: "v2ray-manager.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       022,
	}
}

func runServe(cmd *cobra.Command, args []string) {
	daemon.SetSigHandler(termHandler, syscall.SIGQUIT)
	daemon.SetSigHandler(termHandler, syscall.SIGTERM)
	daemon.SetSigHandler(reloadHandler, syscall.SIGHUP)

	// daemon
	cntxt := buildContext()

	if RunServerDaemon {
		// console.Warning("程序将以daemon形式启动...")

		d, err := cntxt.Reborn()
		if err != nil {
			panic(err)
		}
		// logger.LogIf(err)
		if d != nil {
			return
		}
		defer cntxt.Release()
	}

	// TODO 运行grpc服务
	if RunServerDaemon {
		go runGRPC()
	} else {
		runGRPC()
	}

	if RunServerDaemon {
		err := daemon.ServeSignals()
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
	} else {
		// TODO 暂时不需要注册信号
	}

	log.Println("daemon terminated")
}

func runGRPC() {
	port := config.GetString("app.port")

	// 监听本地端口
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))

	// 注册服务
	manager.RegisterServeServiceServer(s, &serveController.ServeServer{})

	reflection.Register(s)

	logger.Logger.Warn("listening :"+port+"...")
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	log.Println("configuration reloaded")
	return nil
}
