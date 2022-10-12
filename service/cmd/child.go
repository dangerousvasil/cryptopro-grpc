package cmd

import (
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"cryptopro-jsonrpc/src/pool_child"
	"cryptopro-jsonrpc/src/server_internal"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	rootCmd.AddCommand(childCmd)
}

var childCmd = &cobra.Command{
	Use:   "child",
	Short: "A crypto pro child service",
	Long:  `Программное обеспечение, необходимое для работы с электронной подписью`,
	Run: func(cmd *cobra.Command, args []string) {

		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		server := server_internal.NewServiceServer()
		grpc_service.RegisterServiceInternalServer(grpcServer, server)

		_, err = fmt.Fprintln(os.Stdout, pool_child.SERVICE_HI_MSG)
		if err != nil {
			log.Println(err)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
