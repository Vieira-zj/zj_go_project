#!/bin/bash
set -ex
# -x, show run commands with arguments
# -e, tell bash exit script if any statement returns a non-true value
echo "gorun.sh"

# ENV VAR SET
# source $QBOXROOT/kodo/env.sh
# source $QBOXROOT/base/env.sh
ZJ_GOPRJ="${HOME}/Workspaces/zj_go_project"
# if current golang project are not in system path
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
    cd ${ZJ_GOPRJ}/src/tools.test/apps/utilstest;go run main.go
fi

# db test
if [ "$1" == "db" ]; then
    go run src/data.db/main/main.go
fi


# MOCK BIN
function build_bin() {
    target_dir="${HOME}/Downloads/tmp_files"
    target_bin=$1
    cd ${ZJ_GOPRJ}/src/mock.server/main
    GOOS=linux GOARCH=arm go build -o ${target_bin} main.go
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

mock_bin="mockserver"
if [ "$1" == "mock" ]; then
    if [ -z $2 ]; then
        build_bin "${mock_bin}_mac"
    fi
    if [ $2 == "arm" ]; then
        build_bin "${mock_bin}_$2"
    fi
    if [ $2 == "linux" ]; then
        set +e
        target_bin="${mock_bin}_$2"
        build_bin target_bin
        # scp_remote
        set -e
    fi
fi

# build ddtest bin for linux
if [ "$1" == "lxddtest" ]; then
    target_bin="ddtest"
    target_main = "src/tools.test/apps/ddtest/main.go"
    GOOS=linux GOARCH=amd64 go build -o ${target_bin} target_main
    # scp ${target_bin} qboxserver@cs1:~/zhengjin/ && rm ${target_bin}
fi

set +ex # set configs off
