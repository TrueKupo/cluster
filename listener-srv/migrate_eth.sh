#!/usr/bin/env bash

GOOSE_BIN=goose

$GOOSE_BIN -dir ./migrations postgres `cat ./migrationdsn_eth.txt` "$@"
