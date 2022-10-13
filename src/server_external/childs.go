package server_external

import (
	"cryptopro-jsonrpc/src/innchild"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"fmt"
)

func (s *ServerExternal) GetSignChild(req *grpc_service.SignRequest) (*innchild.InnerChild, error) {
	s.lock.RLock()
	child, ok := s.childSignPool[s.GetChildSignKey(req)]
	s.lock.RUnlock()
	if ok {
		return child, nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var err error
	_, ok = s.childSignPool[s.GetChildSignKey(req)]
	if !ok {
		child, err = innchild.NewInSignChild(s.ctx, s.binFile)
		if err != nil {
			return nil, err
		}
		s.childSignPool[s.GetChildSignKey(req)] = child
	}
	return child, err
}

func (s *ServerExternal) GetVerifyChild(req *grpc_service.VerifyRequest) (*innchild.InnerChild, error) {
	s.lock.RLock()
	child, ok := s.childVerifyPool[s.GetChildVerifyKey(req)]
	s.lock.RUnlock()
	if ok {
		return child, nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var err error
	_, ok = s.childVerifyPool[s.GetChildVerifyKey(req)]
	if !ok {
		child, err = innchild.NewInSignChild(s.ctx, s.binFile)
		if err != nil {
			return nil, err
		}
		s.childVerifyPool[s.GetChildVerifyKey(req)] = child
	}
	return child, err
}

func (s *ServerExternal) GetTypeChild() (*innchild.InnerChild, error) {
	if s.childType == nil {
		s.lock.Lock()
		defer s.lock.Unlock()
		if s.childType == nil {
			child, err := innchild.NewInSignChild(s.ctx, s.binFile)
			if err != nil {
				return nil, err
			}
			s.childType = child
		}
	}
	return s.childType, nil
}

func (s *ServerExternal) GetChildSignKey(req *grpc_service.SignRequest) string {
	return fmt.Sprintf(`%s-%s`, req.GetStorage(), req.GetKey())
}

func (s *ServerExternal) GetChildVerifyKey(req *grpc_service.VerifyRequest) string {
	return fmt.Sprintf(`%s`, req.GetStorage())
}
