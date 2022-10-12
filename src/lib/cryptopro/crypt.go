package cryptopro

/*
#cgo CFLAGS: -DUNIX -DHAVE_LIMITS_H -DSIZEOF_VOID_P=8 -I/opt/cprocsp/include/ -I/opt/cprocsp/include/cpcsp -I/opt/cprocsp/include/pki
#cgo LDFLAGS: -L/opt/cprocsp/lib/amd64 -lcapi20 -lcapi10
#include <stdlib.h>
#include <stdarg.h>
#include <CSP_WinCrypt.h>
#include <WinCryptEx.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

const (
	CALG_G28147 = C.CALG_G28147
)

const (
	CRYPT_EXPORTABLE     = C.CRYPT_EXPORTABLE
	CRYPT_USER_PROTECTED = C.CRYPT_USER_PROTECTED
	CRYPT_VERIFYCONTEXT  = C.CRYPT_VERIFYCONTEXT
)

type CryptoProv struct {
	hCryptoProv *C.HCRYPTPROV
	KeySpec     C.uint
}

func CryptImportPublicKeyInfoEx(prov *CryptoProv, context *CertContext) (*Key, error) {
	var pubKey C.HCRYPTKEY

	status := C.CryptImportPublicKeyInfoEx(*prov.hCryptoProv, X509_ASN_ENCODING|PKCS_7_ASN_ENCODING,
		&(*context.pCertContext).pCertInfo.SubjectPublicKeyInfo, 0, 0, nil, &pubKey)
	if status == 0 {
		return nil, errors.New("can't get public key from cert")
	}
	return &Key{hCryptKey: &pubKey}, nil
}

func CryptAcquireContext(container string, provName string, provType uint, flags uint) (*CryptoProv, error) {
	var hProv C.HCRYPTPROV

	var cont *C.char = nil
	if container != "" {
		cont = C.CString(container)
	}

	var prov *C.char = nil
	if provName != "" {
		prov = C.CString(provName)
	}

	status := C.CryptAcquireContext(&hProv, cont, prov, C.uint(provType), C.uint(flags))
	if status == 0 {
		return nil, fmt.Errorf("can't acquire context got error %s", GetLastError())
	}

	return &CryptoProv{hCryptoProv: &hProv}, nil
}
