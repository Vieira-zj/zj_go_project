#!/bin/bash
set -ex
# -x, show run commands with arguments
# -e, tell bash exit script if any statement returns a non-true value
echo "gorun.sh"

# ENV VAR SET
# source $QBOXROOT/kodo/env.sh
# source $QBOXROOT/base/env.sh
ZJ_GOPRJ="${HOME}/Workspaces/zj_go_project"
# if current golang project is not in system path
# GOPATH=${ZJ_GOPRJ}:${GOPATH}


# GO MAIN
# Go learn doc: https://github.com/gopl-zh/gopl-zh.github.com.git
# Go fmt: https://github.com/golang/go/wiki/CodeReviewComments
# Effective Go: https://golang.org/doc/effective_go.html
if [ -z $1 ]; then
    go run src/demo.hello/main/main.go
fi

if [ "$1" == "main" ]; then
    go run src/demo.hello/main/main.go -args hello world
    # go run src/demo.hello/main/main.go -period 3s
    # go run src/demo.hello/main/main.go -h
    # go run src/demo.hello/main/main.go -p 7890 -c 404
fi

if [ "$1" == "util" ]; then
    cd ${ZJ_GOPRJ}/src/tools.app/apps/utilstest;go run main.go
fi

# app test
if [ "$1" == "app" ]; then
    go run src/sys.app/main/main.go
fi


# BUILD MOCK BIN
function go_build_bin() {
    target_dir="${HOME}/Downloads/tmp_files"
    target_bin=$1
    cd ${ZJ_GOPRJ}/src/mock.server/main
    if [ $2 ]; then
        GOOS=linux GOARCH=$2 go build -o ${target_bin} main.go
    else
        go build -o ${target_bin} main.go
    fi
    mv ${target_bin} ${target_dir}
    cp mock_conf.json ${target_dir}
}

function scp_remote() {
    remote_ip="10.200.20.21"
    ping $remote_ip -c 1
    if [ $? == 0 ]; then
        cd ${ZJ_GOPRJ}/src/mock.server/main
        scp ${target_bin} qboxserver@${remote_ip}:~/zhengjin/ && rm ${target_bin}
    fi
}

function build_mock_bin {
    if [[ $1 == "linux" ]]; then
        go_build_bin "${mock_bin}_$1" "amd64"
        # scp_remote
        return
    fi
    if [[ $1 == "arm" ]]; then
        go_build_bin "${mock_bin}_$1" "arm"
        return
    fi
    go_build_bin "${mock_bin}_mac"
}

mock_bin="mockserver"
if [[ $1 == "mock" ]]; then
    build_mock_bin $2
fi


# BUILD DDTEST BIN
function build_ddtest_bin() {
    target_bin="ddtest"
    target_main="src/tools.app/apps/ddtest/main.go"
    GOOS=linux GOARCH=amd64 go build -o ${target_bin} target_main
    # scp ${target_bin} qboxserver@cs1:~/zhengjin/ && rm ${target_bin}
}

if [[ $1 == "ddtest" ]]; then
    build_ddtest_bin
fi

echo "go build and run DONE."
set +ex # set configs off
