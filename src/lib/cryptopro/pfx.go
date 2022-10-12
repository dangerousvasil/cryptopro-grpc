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
	"os"
	"unicode/utf16"
	"unsafe"
)

type CryptDataBlob struct {
	DataBlob C.CRYPT_DATA_BLOB
}

func PFXImportCertStore(pfxPath, password string) (*CertStore, error) {
	var status C.int

	pfxBlob, err := os.ReadFile(pfxPath)
	if err != nil {
		return nil, err
	}

	var crypApiBlob = CryptDataBlob{}

	test := C.CBytes(pfxBlob)

	crypApiBlob.DataBlob.pbData = (*C.uchar)(test)
	crypApiBlob.DataBlob.cbData = C.uint(len(pfxBlob))

	status, err = C.PFXIsPFXBlob(&crypApiBlob.DataBlob)
	if err != nil {
		return nil, GetLastError()
	}
	if status == 0 {

		return nil, GetLastError()
	}

	wstr := utf16.Encode([]rune(password))
	wstr = append(wstr, 0x00)

	dwImportFlags := CRYPT_EXPORTABLE | CRYPT_USER_PROTECTED

	store, err := C.PFXImportCertStore(&crypApiBlob.DataBlob, (C.LPCWSTR)(unsafe.Pointer(&wstr[0])), C.uint(dwImportFlags))
	if err != nil {
		return nil, GetLastError()
	}

	return &CertStore{HCertStore: &store}, nil
}
