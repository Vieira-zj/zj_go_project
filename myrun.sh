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



# check field exist
# export is_exist=y
# if [[ "${is_exist}" ]]; then
#     echo "is exist"
# fi

# check array
# tmp_list1=("ele1")
# tmp_list2=("ele2" "ele3")
# tmp_list3=(${tmp_list1[@]} ${tmp_list2[@]})
# for ele in ${tmp_list3[@]}; do
#     echo $ele
# done
# echo ${tmp_list3[@]}
# echo "length: ${#tmp_list3[@]}"

# run download parallel, wait and stop
for (( i=0; i<20; i++)); do
echo "run at: $i"
curl -v "http://p1vvj0112.bkt.clouddn.com/512M" -x iovip-z2.qbox.me:80 > /dev/null &
done
sleep 3m
ps -ef | grep "curl" | grep -v "grep" | awk '{print $2}' | xargs kill -9

# custom functions
# echoEnv() { echo "TEST_ENV=$TEST_ENV"; echo "TEST_ZONE=$TEST_ZONE";}
# setEnv() { export TEST_ENV=$1; echo "TEST_ENV=$TEST_ENV";}
# setZone() { export TEST_ZONE=$1; echo "TEST_ZONE=$TEST_ZONE";}
# run function
# echoEnv

set +x # set config x off
