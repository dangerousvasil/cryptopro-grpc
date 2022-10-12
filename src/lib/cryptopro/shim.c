#include <stdlib.h>
#include <stdarg.h>
#include <string.h>
#include <CSP_WinCrypt.h>

#include "shim.h"
#include "_cgo_export.h"

CERT_BLOB* get_blob(PCCERT_CONTEXT cert) {
	CERT_BLOB* blob = malloc(sizeof(CERT_BLOB));

	blob->cbData = cert->cbCertEncoded;
	blob->pbData = cert->pbCertEncoded;

	return blob;
}


CMSG_SIGNER_ENCODE_INFO* init_signer(PCERT_INFO cert_info, HCRYPTPROV h_crypt_prov, char* hash_algo) {
	CMSG_SIGNER_ENCODE_INFO* signer;
	CRYPT_ALGORITHM_IDENTIFIER *hash_ident;

	hash_ident = malloc(sizeof(CRYPT_ALGORITHM_IDENTIFIER));
	memset(hash_ident, 0, sizeof(CRYPT_ALGORITHM_IDENTIFIER));
	hash_ident->pszObjId = hash_algo;

	signer = malloc(sizeof(CMSG_SIGNER_ENCODE_INFO));
	memset(signer, 0, sizeof(CMSG_SIGNER_ENCODE_INFO));
	signer->cbSize = sizeof(CMSG_SIGNER_ENCODE_INFO);
	signer->pCertInfo = cert_info;
	signer->hCryptProv = h_crypt_prov;
	signer->HashAlgorithm = *hash_ident;
	signer->dwKeySpec = AT_KEYEXCHANGE;
	signer->pvHashAuxInfo = NULL;

	return signer;
}

CERT_EXTENSION* get_extension(PCERT_INFO cert_info, int index) {
    return &cert_info->rgExtension[index];
}

PCERT_ALT_NAME_INFO get_dist_point(PCRL_DIST_POINTS_INFO dist_info, int index) {
    return &dist_info->rgDistPoint[index].DistPointName._empty_union_.FullName;
}

LPWSTR get_dist_point_url(PCERT_ALT_NAME_INFO nameInfo, int index) {
    return nameInfo->rgAltEntry[index]._empty_union_.pwszURL;
}

LPSTR get_access_method(PCERT_AUTHORITY_INFO_ACCESS info_access, int index) {
    return info_access->rgAccDescr[index].pszAccessMethod;
}

LPWSTR get_access_location(PCERT_AUTHORITY_INFO_ACCESS info_access, int index) {
    return info_access->rgAccDescr[index].AccessLocation._empty_union_.pwszURL;
}
