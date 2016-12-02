#!/bin/bash

set -e -x

go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get github.com/gorilla/mux
ginkgo -r
