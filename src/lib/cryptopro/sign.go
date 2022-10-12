package cryptopro

/*
#cgo CFLAGS: -DUNIX -DHAVE_LIMITS_H -DSIZEOF_VOID_P=8 -I/opt/cprocsp/include/ -I/opt/cprocsp/include/cpcsp -I/opt/cprocsp/include/pki
#cgo LDFLAGS: -L/opt/cprocsp/lib/amd64 -lcapi20 -lcapi10 -lrdrsup -lssp
#include <stdlib.h>
#include <stdarg.h>
#include <CSP_WinCrypt.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
)

func CryptSignHash(hHash *CryptoHash, keySpec uint, flags uint) ([]byte, error) {
	var SigLen C.uint

	if hHash == nil {
		return nil, errors.New("hHash can't be nil")
	}

	status := C.CryptSignHash(*hHash.hHash, C.uint(keySpec), nil, C.uint(flags), nil, &SigLen)
	if status == 0 {
		log.Println(GetLastError())
		return nil, fmt.Errorf("can't sign hash size got eror %s", GetLastError())
	}

	signature := make([]byte, SigLen)
	status = C.CryptSignHash(*hHash.hHash, C.uint(keySpec), nil, C.uint(flags), (*C.uchar)(&signature[0]), &SigLen)
	if status == 0 {
		return nil, fmt.Errorf("can't sign hash got error %s", GetLastError())
	}

	return signature, nil
}
