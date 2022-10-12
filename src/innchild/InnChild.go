package innchild

import (
	"context"
	"cryptopro-jsonrpc/src/lib/free_port"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os/exec"
)

type InnerChild struct {
	child          *exec.Cmd
	InternalClient grpc_service.ServiceInternalClient
	port           int
}

func NewInSignChild(ctx context.Context, binFile string) (*InnerChild, error) {
	var err error
	port, err := free_port.GetFreePort()
	if err != nil {
		return nil, err
	}
	o := InnerChild{
		port: port,
	}

	o.child, err = StartChild(binFile, port)
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(`localhost:%d`, o.port), opts...)
	if err != nil {
		log.Println(`Dial`, err)
		return nil, err
	}
	o.InternalClient = grpc_service.NewServiceInternalClient(conn)

	return &o, nil
}
