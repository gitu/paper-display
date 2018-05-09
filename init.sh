#!/bin/bash

set -e

export GOPATH=/home/pi/go/
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin

go get -u github.com/gitu/paper-display

paper-display
