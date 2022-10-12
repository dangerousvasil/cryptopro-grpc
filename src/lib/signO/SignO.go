package signo

import (
	"cryptopro-jsonrpc/src/lib/cryptopro"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"fmt"
	"log"
)

type SignO struct {
	Key     string
	Storage string

	Store  *cryptopro.CertStore
	Client *cryptopro.CertContext
}

var Obj *SignO

func SignOGo(req *grpc_service.SignRequest) ([]byte, error) {

	if Obj == nil {
		Obj = new(SignO)
		Obj.Key = req.Key
		Obj.Storage = req.Storage
	}

	var err error

	if Obj.Store == nil {
		Obj.Store, err = cryptopro.CertOpenSystemStore(req.Storage)
		//defer cryptopro.CertCloseStore(store, cryptopro.CERT_CLOSE_STORE_CHECK_FLAG)
		if err != nil {
			log.Println(fmt.Sprintf(`open store: %s (%s)`, req.Storage, err))
			return nil, err
		}
	}

	if Obj.Client == nil {
		Obj.Client, err = cryptopro.CertFindCertificateInStore(Obj.Store, req.Key, cryptopro.CERT_FIND_SHA1_HASH)
		//defer cryptopro.CertFreeCertificateContext(client)
		if err != nil {
			log.Println(fmt.Sprintf(`open cert: %s (%s)`, req.Key, err))
			return nil, err
		}
	}

	msgInfo, err := cryptopro.InitSignedInfo(Obj.Client)
	if err != nil {
		log.Println(`InitSignedInfo :` + err.Error())
		return nil, err
	}

	msg, err := cryptopro.CryptMsgOpenToEncode(msgInfo, cryptopro.CMSG_SIGNED, 0, nil)
	if err != nil {
		return nil, err
	}
	defer cryptopro.CryptMsgClose(msg)

	errUpd := cryptopro.CryptMsgUpdate(msg, req.Content, 1)
	if errUpd != nil {
		return nil, err
	}

	content, err := cryptopro.CryptMsgGetParam(msg, cryptopro.CMSG_CONTENT_PARAM, 1)
	if err != nil {
		return nil, err
	}

	return content, err
}
