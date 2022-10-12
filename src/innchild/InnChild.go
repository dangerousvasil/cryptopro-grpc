package innchild

import (
	"context"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os/exec"
)

type InnerChild struct {
	child          *exec.Cmd
	InternalClient grpc_service.ServiceInternalClient
}

func NewInSignChild(ctx context.Context, binFile string) (*InnerChild, error) {
	var err error
	o := InnerChild{}

	o.child, err = StartChild(binFile)
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.DialContext(ctx, `localhost:8081`, opts...)
	if err != nil {
		log.Println(`Dial`, err)
		return nil, err
	}
	o.InternalClient = grpc_service.NewServiceInternalClient(conn)

	return &o, nil
}
