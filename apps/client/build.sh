#!/bin/bash
CGO_ENABLED=0 go build -a -installsuffix cgo -o /app/dist/client /app/apps/client/src