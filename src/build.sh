#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=migrateDockerRegistries
fi

go build -o /opt/bin/${BINARY} .
