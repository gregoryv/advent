#!/bin/bash -e

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

if [ -e ci.sh ]; then
    set -o xtrace
    ./ci.sh
    exit
fi

case $extension in
    go)
	set -o xtrace	
	export GOPROXY="https://goproxy.io"
        goimports -w $path
        ;;
esac

go test ./...

