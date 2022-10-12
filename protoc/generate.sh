#!/bin/bash
FILE=$( realpath ${BASH_SOURCE[0]} )
DIR="`dirname ${FILE}`"
##=-=-=-=-=-=-=-=-=-=-=##
echo "Clean"
rm -f ${DIR}/../src/lib/grpc_service/*.go
echo "Generate"
mkdir -p ${DIR}/../src/lib/grpc_service/
protoc --proto_path=${DIR} --go_out=${DIR}/../src/lib/grpc_service/ --go-grpc_out=${DIR}/../src/lib/grpc_service/ ${DIR}/*.proto
echo "Done"