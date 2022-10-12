package server_external

import (
	"context"
	"cryptopro-jsonrpc/src/innchild"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type ServerExternal struct {
	grpc_service.ServiceServer
	binFile        string
	child          *exec.Cmd
	serviceStream  grpc_service.ServiceInternal_SignClient
	InternalClient grpc_service.ServiceInternalClient

	childPool map[string]*innchild.InnerChild
	lock      *sync.RWMutex
	ctx       context.Context
}

func NewServiceServer(ctx context.Context, binFile string) *ServerExternal {
	return &(ServerExternal{
		ctx:       ctx,
		binFile:   binFile,
		lock:      new(sync.RWMutex),
		childPool: map[string]*innchild.InnerChild{},
	})
}

func (s *ServerExternal) Sign(ctx context.Context, req *grpc_service.SignRequest) (*grpc_service.SignResponse, error) {
	var err error
	ctx = context.Background()

	child, err := s.GetChild(req)
	if err != nil {
		log.Println(`GetChild`, err)
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

func (s *ServerExternal) GetChild(req *grpc_service.SignRequest) (*innchild.InnerChild, error) {
	s.lock.RLock()
	child, ok := s.childPool[s.GetChildKey(req)]
	s.lock.RUnlock()
	if ok {
		return child, nil
	}
	s.lock.Lock()
	var err error
	defer s.lock.Unlock()
	s.childPool[s.GetChildKey(req)], err = innchild.NewInSignChild(s.ctx, s.binFile)
	return s.childPool[s.GetChildKey(req)], err
}

func (s *ServerExternal) GetChildKey(req *grpc_service.SignRequest) string {
	return fmt.Sprintf(`%s-%s`, req.GetStorage(), req.GetKey())
}
