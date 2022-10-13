package server_internal

import (
	"cryptopro-jsonrpc/src/lib/cryptopro"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	signo "cryptopro-jsonrpc/src/lib/signO"
	"unsafe"
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

func (s *ServerInternal) Type(stream grpc_service.ServiceInternal_TypeServer) error {
	data, err := stream.Recv()
	if err != nil {
		return err
	}

	resp := grpc_service.TypeResponse{
		Type:        uint32(len(data.Content)),
		Description: `this method has no implementation`,
	}

	err = stream.Send(&resp)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerInternal) Verify(stream grpc_service.ServiceInternal_VerifyServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	resp := grpc_service.VerifyResponse{}

	decMsg, err := cryptopro.CryptMsgOpenToDecode(0, cryptopro.CMSG_DETACHED_FLAG, nil)
	if err != nil {
		return err
	}
	defer cryptopro.CryptMsgClose(decMsg)

	err = cryptopro.CryptMsgUpdate(decMsg, req.Content, 1)
	if err != nil {
		return err
	}

	data, err := cryptopro.CryptMsgGetParam(decMsg, cryptopro.CMSG_SIGNER_CERT_INFO_PARAM, 0)
	if err != nil {
		return err
	}

	if s.store == nil {
		s.store, err = cryptopro.CertOpenSystemStore(req.GetStorage())
		if err != nil {
			return err
		}
	}

	checkCert, err := cryptopro.CertGetSubjectCertificateFromStore(s.store, data)
	if err != nil {
		return err
	}
	defer cryptopro.CertFreeCertificateContext(checkCert)

	status, err := cryptopro.CryptMsgControl(decMsg, 0, cryptopro.CMSG_CTRL_VERIFY_SIGNATURE, unsafe.Pointer(checkCert.GetCertInfo()))
	if err != nil {
		return err
	}

	if status {
		resp.Key = checkCert.GetThumbprint()
	}

	err = stream.Send(&resp)
	if err != nil {
		return err
	}
	return nil
}
