#!/bin/bash
echo "Debug started..."
dlv debug main.go
echo "Debug ended."
go build -o myprogram main.go 
dlv exec myprogram
