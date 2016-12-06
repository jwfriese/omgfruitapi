#!/bin/bash

set -e -x

OUTPUT_DIR=$(pwd)/new-release
echo $OUTPUT_DIR

cd $GOPATH/src/github.com/jwfriese/omgfruitapi

git config --global user.name "Jared Friese"
git config --global user.email "jared.friese@gmail.com"

go build main.go
mv main $OUTPUT_DIR/release-binary

NEXT=$(date +%s)

echo "$NEXT" > $OUTPUT_DIR/name
echo "$NEXT" > $OUTPUT_DIR/tag
echo "Version $NEXT" > $OUTPUT_DIR/body

