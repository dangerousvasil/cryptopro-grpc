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

var CRYPT_E_STREAM_MSG_NOT_READY = errors.New(`0x80091010`)
var CRYPT_E_NOT_FOUND = errors.New(`0x80092004`)
var CRYPT_E_INVALID_ALG = errors.New(`0x80090008`)

type CryptoproError struct {
	CodeError int
}

func (e CryptoproError) Code() int {
	return e.CodeError
}

func (e CryptoproError) Error() string {
	return fmt.Sprintf("0x%x", e.CodeError)
}

func GetLastError() *CryptoproError {
	e := new(CryptoproError)
	e.CodeError = int(C.GetLastError())
	return e
}
