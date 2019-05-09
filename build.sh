#!/usr/bin/env bash
export GOPROXY=https://goproxy.io
go build
sudo supervisorctl reload