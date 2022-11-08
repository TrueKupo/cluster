#!/bin/sh

MODULE_PWD=$PWD
SERVICES_PWD=interfaces/services
MESSAGES_PWD=interfaces/messages

cd ${MODULE_PWD}/${MESSAGES_PWD}

for d in *; do
    if ! [ -d $d ]; then
        continue
    fi

    if [ -f ./$d/$d.proto ]; then
        protoc -I. \
					-I../ \
					--go_out=. \
					--go-grpc_out=. \
					--go_opt=paths=source_relative \
					$d/$d.proto
        echo "messages: $d.proto"
    fi
done

cd ${MODULE_PWD}/${SERVICES_PWD}

for d in *; do
    if ! [ -d $d ]; then
        continue
    fi

    if [ -f ./$d/$d.proto ]; then
        protoc -I. \
					-I../ \
					--go_out=. \
					--go-grpc_out=. \
					--go_opt=paths=source_relative \
					$d/$d.proto
        echo "services: $d.proto"
    fi
done
