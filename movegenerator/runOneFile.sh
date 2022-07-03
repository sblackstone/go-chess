#!/usr/bin/env bash

go test -v $1 `ls *.go|grep -v _test`
