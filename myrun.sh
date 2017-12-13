#!/bin/bash
set -x # show run commands with arguments

# main
# go run src/demo.hello/main/main.go
# go run src/demo.tests/main/test.go

# go tests, base $GOPATH
# go test -v demo.tests/gotests/
# go test -v src/demo.tests/gotests/word_test.go

# bdd tests
ginkgo -v -focus="describe" src/demo.tests/bddtests/


set +x # set config x off
