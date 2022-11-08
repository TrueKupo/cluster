#!/usr/bin/env bash

GOOSE_BIN=/home/vvv/goose/goose

$GOOSE_BIN -dir ./migrations postgres `cat ./migrationdsn.txt` "$@"
