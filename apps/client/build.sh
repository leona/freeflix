#!/bin/bash
CGO_ENABLED=0 go build -a -installsuffix cgo -o ../../dist/client ./src