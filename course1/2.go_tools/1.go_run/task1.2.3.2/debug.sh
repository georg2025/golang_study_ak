#!/bin/bash
echo "Debug started..."
dlv debug "$1"
echo "Debug ended."
go build -o myprogram main.go 
dlv exec myprogram
