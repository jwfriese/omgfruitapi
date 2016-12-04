#!/bin/bash

set -e

ls
ls /
OUTPUT_DIR=$(pwd)/new-release
echo $OUTPUT_DIR

cd $GOPATH/src/github.com/jwfriese/omgfruitapi

go build main.go
mv main $OUTPUT_DIR/release-binary

NEXT=$(cat /proc/sys/kernel/random/uuid)

echo "$NEXT" > $OUTPUT_DIR/name
echo "$NEXT" > $OUTPUT_DIR/tag
echo "Version $NEXT" > $OUTPUT_DIR/body

