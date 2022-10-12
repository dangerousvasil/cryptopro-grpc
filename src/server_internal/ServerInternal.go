package server_internal

import (
	"cryptopro-jsonrpc/src/lib/cryptopro"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	signo "cryptopro-jsonrpc/src/lib/signO"
)

type ServerInternal struct {
	grpc_service.ServiceInternalServer

	// env for Sign process
	store  *cryptopro.CertStore
	client *cryptopro.CertContext
}

func NewServiceServer() *ServerInternal {
	return new(ServerInternal)
}

func (s *ServerInternal) Sign(stream grpc_service.ServiceInternal_SignServer) error {
	data, err := stream.Recv()
	if err != nil {
		return err
	}

	resp := grpc_service.SignResponse{}
	resp.Content, err = signo.SignOGo(data)
	if err != nil {
		return err
	}

	err = stream.Send(&resp)
	if err != nil {
		return err
	}
	return nil
}
