#!/bin/bash

GOPATH="/home/pi/go/"
EXPORT GOPATH

go get -u github.com/gitu/paper-display

/home/pi/go/bin/paper-display
