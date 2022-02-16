#!/usr/bin/env bash
go test -coverprofile blarg ./...
go tool cover -html blarg
rm -f blarg
