#!/bin/bash

set -e -x

go get github.com/onsi/ginkgo
ginkgo -r
