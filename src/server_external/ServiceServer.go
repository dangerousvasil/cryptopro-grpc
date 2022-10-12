package server_external

import (
	"context"
	"cryptopro-jsonrpc/src/innchild"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"fmt"
	"log"
	"sync"
)

type ServerExternal struct {
	grpc_service.ServiceServer
	binFile string

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

const CprocspDefaultStorage = `my`

func (s *ServerExternal) Sign(ctx context.Context, req *grpc_service.SignRequest) (*grpc_service.SignResponse, error) {
	var err error
	ctx = context.Background()

	if req.GetStorage() == `` {
		req.Storage = CprocspDefaultStorage
	}

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
	defer s.lock.Unlock()
	var err error
	_, ok = s.childPool[s.GetChildKey(req)]
	if !ok {
		child, err = innchild.NewInSignChild(s.ctx, s.binFile)
		if err != nil {
			return nil, err
		}
		s.childPool[s.GetChildKey(req)] = child
	}
	return child, err
}

func (s *ServerExternal) GetChildKey(req *grpc_service.SignRequest) string {
	return fmt.Sprintf(`%s-%s`, req.GetStorage(), req.GetKey())
}
