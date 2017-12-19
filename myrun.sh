#!/bin/bash
set -x # show run commands with arguments

# main
go run src/demo.hello/main/main.go
# go run src/demo.hello/main/main.go hello world
# go run src/demo.hello/main/main.go -period 3s
# go run src/demo.tests/main/test.go

# go tests, base $GOPATH
# go test -v demo.tests/gotests/
# go test -v demo.tests/gotests/ -run="TestEcho"
# go test -v src/demo.tests/gotests/word_test.go

# bdd tests
# ginkgo -v -focus="demo01" src/demo.tests/bddtests/
# ginkgo -v -focus="describe table" src/demo.tests/bddtests/
# ginkgo -v -focus="parallel" src/demo.tests/bddtests/


set +x # set config x off
