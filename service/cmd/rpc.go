package cmd

import (
	"context"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"cryptopro-jsonrpc/src/server_external"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rpcPort *int

func init() {
	rootCmd.AddCommand(rpcCmd)
	rpcPort = rpcCmd.Flags().IntP(`port`, `p`, 8080, `port run on`)
}

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start json rpc service",
	Long:  `Сервис обработки json rpc запросов`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *rpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)

		server := server_external.NewServiceServer(ctx, binFile)

		grpc_service.RegisterServiceServer(grpcServer, server)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			for {
				switch <-sigChan {
				case syscall.SIGTERM:
					log.Println("Recive signal SIGTERM")
					cancel()
					grpcServer.GracefulStop()
				case syscall.SIGINT:
					log.Println("Recive signal SIGINT")
					cancel()
					grpcServer.GracefulStop()
				default:
					time.Sleep(time.Millisecond * 100)
				}
			}
		}()

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Panicln(err)
		}

	},
}
