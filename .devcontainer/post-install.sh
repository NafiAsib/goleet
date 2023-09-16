#!/bin/bash
echo "Installing go watch"
go install github.com/mitranim/gow@latest
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"