#!/bin/bash

set -e

go build $GOPATH/src/github.com/jwfriese/omgfruitapi/main.go
mv main new-release/release-binary

CURRENT=$(git tag --sort="-refname" | head -1)
NEXT="$(($CURRENT+1))"

echo "$NEXT" > new-release/name
echo "$NEXT" > new-release/tag
echo "Version $NEXT" > new-release/body

git tag -m "Version $NEXT" $NEXT
git push --tags

