#!/bin/bash

if test "$1" == "" ; then
	echo $'\a'Supply the name of a proto, with the directory structure named after it, e.g:
	echo $0 greet
	exit
fi

SERVICE=$1
PBPATH="${SERVICE}/${SERVICE}pb/${SERVICE}.proto"
echo $PBPATH
protoc $PBPATH --go_out=plugins=grpc:.
