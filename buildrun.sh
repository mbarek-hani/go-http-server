#!/bin/sh
pkill http-server && echo "Sent kill"
rm -f ./http-server
clear
go build && ./http-server