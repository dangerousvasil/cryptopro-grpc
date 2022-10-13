package cmd

import (
	"cryptopro-jsonrpc/src/innchild"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"cryptopro-jsonrpc/src/server_internal"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var childPort *int

func init() {
	rootCmd.AddCommand(childCmd)
	childPort = childCmd.Flags().IntP(`port`, `p`, 8080, `port run on`)
}

var childCmd = &cobra.Command{
	Use:   "child",
	Short: "A crypto pro child service",
	Long:  `Программное обеспечение, необходимое для работы с электронной подписью`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *childPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		server := server_internal.NewServiceServer()
		grpc_service.RegisterServiceInternalServer(grpcServer, server)

		_, err = fmt.Fprintln(os.Stdout, innchild.SERVICE_MSG)
		if err != nil {
			log.Println(err)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
