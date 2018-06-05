#!/bin/bash
set -x -e
# -x, show run commands with arguments
# -e, tell bash exit script if any statement returns a non-true value
echo "myrun.sh"

# ENV VAR SET
# source $QBOXROOT/kodo/env.sh
# source $QBOXROOT/base/env.sh
ZJGOPRJ="${HOME}/Workspaces/zj_projects/ZjGoProject"
GOPATH=${ZJGOPRJ}:${GOPATH}


# MAIN
# demo test
go run src/demo.hello/main/main.go
if [[ "$1" == "main" ]]; then
    go run src/demo.hello/main/main.go hello world
    # go run src/demo.hello/main/main.go -h
    # go run src/demo.hello/main/main.go -p 7890 -c 404
fi

# db test
if [[ "$1" == "db" ]]; then
    go run src/data.db/main/main.go
fi


# BIN
# build bin
if [[ "$1" == "bin" ]]; then
    go build -o mockserver src/mock.server/main/main.go
fi

# build bin for linux
if [[ "$1" == "lxbin" ]]; then
    target_bin="mockserver"
    GOOS=linux GOARCH=amd64 go build -o ${target_bin} src/mock.server/main/main.go
    scp ${target_bin} qboxserver@10.200.20.21:~/zhengjin/ && rm ${target_bin}
fi


# GO TEST
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


# BDD TEST
# bddtest="/Users/zhengjin/gopath/bin/ginkgo"
# ${bddtest} -v -focus="test.hooks" src/demo.tests/bddtests/

# ginkgo -v -focus="demo01.routine.done" src/demo.tests/bddtests/
# ginkgo -v -focus="demo02.share" src/demo.tests/bddtests/


# SHELL EXAMPLES
# EX01, field check
# https://blog.csdn.net/longyinyushi/article/details/50728049

# EX01-01, field exist check
# is_exist="test"
# if [ ${is_exist} ]; then
#     echo "[] check, is exist."
# fi
# if [[ "${is_exist}" ]]; then
#     echo "[[]] check, is exist."
# fi

# EX01-02, field length check
# empty_str=""
# if [ -z $empty_str ]; then
#     echo "string length = 0."
# else
#     echo "string length > 0."
# fi

# EX01-03, field empty check
# test_str="test"
# if [[ -n $test_str ]]; then
#     echo 'string not empty.'
# else
#     echo 'string empty.'
# fi


# EX02-01, array
# tmp_list1=("ele1")
# tmp_list2=("ele2" "ele3")
# tmp_list3=(${tmp_list1[@]} ${tmp_list2[@]})
# for ele in ${tmp_list3[@]}; do
#     echo $ele
# done
# echo ${tmp_list3[@]} # echo all elements
# echo "length: ${#tmp_list3[@]}"

# EX02-02
# focus_pkg=()
# temp_pkg1=("a1" "a2")
# focus_pkg=(${focus_pkg[@]} ${temp_pkg1[@]}) # append
# echo ${focus_pkg[@]}
# temp_pkg2=("a3" "a4" "a5")
# focus_pkg=(${focus_pkg[@]} ${temp_pkg2[@]})
# echo ${focus_pkg[@]}

# EX02-03
# focus_pkg=("a1" "a2" "a3" "a4")
# skip_pkg="a1,a2,a3,a4,a5,a6,a7"
# for v in ${focus_pkg[@]}; do
#     skip_pkg=${skip_pkg/${v},/} # replace {$v}, with ""
# done
# echo "focus packages => ${focus_pkg[@]}"
# echo "skip packages => ${skip_pkg}"


# EX03, if-else with regexp
# node_name="go1.9_fix"
# if [[ ($node_name =~ "go1.9") && ($node_name =~ "fix") ]]; then
#     echo 'version check ok.'
# else
#     echo 'version should be go1.9 with fix.'
# fi


# EX04, ${var} usage
# tmp_file=/dir1/dir2/dir3/my.file.txt
# # sub string start 0, len = 5
# echo ${tmp_file:0:5}
# # sub string start 5, len = 5
# echo ${tmp_file:5:5}
# # replace first "dir" with "path"
# echo ${tmp_file/dir/path}
# # replace all "dir" with "path"
# echo ${tmp_file//dir/path}


# EX05, +=
# src="hello"
# src=${src}" world"
# echo ${src}
# src="test, ${src}"
# echo ${src}


# EX06, iterator
# for i in $(seq 1 3); do
#     echo "index ${i}"
# done

# for (( i = 0; i < 3; i++ )); do
#     echo "index ${i}"
# done

# for f in $(ls ~/Downloads/tmp_files/test.*); do
#     echo "test file: ${f}"
# done


# EX07, run download parallel
# for (( i=0; i<20; i++)); do
# echo "run at: $i"
# curl -v "http://7zkl9d.com1.z1.glb.clouddn.com/slowResponse" -x iovip-z1.qbox.me:80 > /dev/null &
# sleep 2s
# done
# sleep 15m
# ps -ef | grep "curl" | grep -v "grep" | awk '{print $2}' | xargs kill -9


# EX08, custom functions
# echoEnv() { echo "TEST_ENV=$TEST_ENV"; echo "TEST_ZONE=$TEST_ZONE";}
# setEnv() { export TEST_ENV=$1; echo "TEST_ENV=$TEST_ENV";}
# setZone() { export TEST_ZONE=$1; echo "TEST_ZONE=$TEST_ZONE";}

# findStr() { grep "$1" ./*;}
# findStrAll() { grep -r "$1" ./;}

# run function
# echoEnv
# findStrAll "search_text"


# EX09, tips
# echo "current dir: $(pwd)"

# echo "test exit with error code 1."
# exit 1


set +x +e # set configs off
