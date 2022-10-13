package server_external

import (
	"context"
	"cryptopro-jsonrpc/src/innchild"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"log"
	"sync"
)

type ServerExternal struct {
	grpc_service.ServiceServer
	binFile string

	childSignPool   map[string]*innchild.InnerChild
	childVerifyPool map[string]*innchild.InnerChild
	childType       *innchild.InnerChild
	lock            *sync.RWMutex
	ctx             context.Context
}

func NewServiceServer(ctx context.Context, binFile string) *ServerExternal {
	return &(ServerExternal{
		ctx:             ctx,
		binFile:         binFile,
		lock:            new(sync.RWMutex),
		childSignPool:   map[string]*innchild.InnerChild{},
		childVerifyPool: map[string]*innchild.InnerChild{},
	})
}

const CprocspDefaultStorage = `my`

func (s *ServerExternal) Type(ctx context.Context, req *grpc_service.ContentRequest) (*grpc_service.TypeResponse, error) {

	child, err := s.GetTypeChild()
	if err != nil {
		log.Println(`GetSignChild`, err)
		return nil, err
	}

	serviceStream, err := child.InternalClient.Type(ctx)
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

func (s *ServerExternal) Sign(ctx context.Context, req *grpc_service.SignRequest) (*grpc_service.SignResponse, error) {
	var err error

	if req.GetStorage() == `` {
		req.Storage = CprocspDefaultStorage
	}

	child, err := s.GetSignChild(req)
	if err != nil {
		log.Println(`GetSignChild`, err)
		return nil, err
	}

	serviceStream, err := child.InternalClient.Sign(ctx)
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

func (s *ServerExternal) Verify(ctx context.Context, req *grpc_service.VerifyRequest) (*grpc_service.VerifyResponse, error) {
	var err error

	if req.GetStorage() == `` {
		req.Storage = CprocspDefaultStorage
	}

	child, err := s.GetVerifyChild(req)
	if err != nil {
		log.Println(`GetSignChild`, err)
		return nil, err
	}

	serviceStream, err := child.InternalClient.Verify(ctx)
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
