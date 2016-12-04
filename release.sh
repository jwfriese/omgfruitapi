#!/bin/bash

set -e

CURRENT=$(git tag --sort="-refname" | head -1)
NEXT="$(($CURRENT+1))"
git tag -m "Version $NEXT" $NEXT
git push --tags
