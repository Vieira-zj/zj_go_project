#!/bin/bash
set -x # show run commands with arguments
echo "here"

# main
go run src/demo.hello/main/main.go
# go run src/demo.hello/main/main.go hello world
# go run src/demo.hello/main/main.go -h
# go run src/demo.hello/main/main.go -p 7890 -c 404


# go tests, root_dir = $GOPATH
# go help test
# go test -v demo.tests/gotests/
# go test -v -run="TestEcho" demo.tests/gotests/
# go test -v src/demo.tests/gotests/word_test.go

# Compile the test binary to pkg.test but do not run it.
# The file name can be changed with the -o flag.
# go test -c

# go tests, coverage
# go test -v -run="IsPalindrome" -cover -coverprofile=c.out demo.tests/gotests/
# go test -v -cover -coverprofile=c.out demo.tests/gotests/
# go tool cover -html=c.out

# go tests, benchmark
# go test -v -bench=. src/demo.tests/gotests/word_ben_test.go
# go test -v -bench=. -benchmem src/demo.tests/gotests/word_ben_test.go


# bdd tests
# ginkgo -v -focus="demo01.routine.done" src/demo.tests/bddtests/
# ginkgo -v -focus="demo02.DescribeTable" src/demo.tests/bddtests/

set +x # set config x off
