package server_external

import (
	"context"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"cryptopro-jsonrpc/src/pool_child"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os/exec"
)

type ServerExternal struct {
	grpc_service.ServiceServer
	binFile        string
	child          *exec.Cmd
	serviceStream  grpc_service.ServiceInternal_SignClient
	InternalClient grpc_service.ServiceInternalClient
}

func NewServiceServer(binFile string) *ServerExternal {
	return &(ServerExternal{
		binFile: binFile,
	})
}

func (s *ServerExternal) Sign(ctx context.Context, req *grpc_service.SignRequest) (*grpc_service.SignResponse, error) {
	var err error
	ctx = context.Background()

	if s.child == nil {

		s.child, err = pool_child.StartChild(s.binFile)
		if err != nil {
			return nil, err
		}

		var opts []grpc.DialOption

		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		opts = append(opts, grpc.WithBlock())

		//defer cancel()
		conn, err := grpc.DialContext(ctx, `localhost:8081`, opts...)
		if err != nil {
			log.Println(`Dial`, err)
			return nil, err
		}
		s.InternalClient = grpc_service.NewServiceInternalClient(conn)
	}
	serviceStream, err := s.InternalClient.Sign(ctx)
	if err != nil {
		log.Println(`Dial`, err)
		return nil, err
	}

	err = serviceStream.Send(req)
	if err != nil {
		log.Println(`Send`, err)
		return nil, err
	}

	return serviceStream.Recv()
}
