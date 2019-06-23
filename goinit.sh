#!/bin/bash
set -ex
echo "goinit.sh"

function add_sub_git_repos() {
    repos=("src/github.com/derekparker/delve" "src/github.com/ianthehat/godef" \
        "src/github.com/karrick/godirwalk" "src/github.com/mdempsky/gocode" \
        "src/github.com/ramya-rao-a/go-outline" "src/github.com/rogpeppe/godef" \
        "src/github.com/sqs/goreturns" "src/github.com/stamblerre/gocode" \
        "src/github.com/uudashr/gopkgs"
    )
    
    cd ${HOME}/Workspaces/zj_go_project
    for repo in ${repos[*]}; do
        git rm --cached -f ${repo}
        git submodule add "https://${repo:4}" ${repo}
    done
}

# install external tools for vscode golang
function install_ext_tools() {
    ext_tools=("github.com/mdempsky/gocode" "github.com/ramya-rao-a/go-outline" \
        "github.com/acroca/go-symbols" "golang.org/x/tools/cmd/guru" \
        "golang.org/x/tools/cmd/gorename" "github.com/stamblerre/gocode" \
        "github.com/sqs/goreturns" "golang.org/x/lint/golint"
    )
    
    for tool in ${ext_tools[*]}; do
        echo "install tool: ${tool}"
        cd ${HOME}/Workspaces/zj_go_project/src
        cd tool;go install
    done
}

# main
# install_ext_tools
# fix_sub_repo

set +ex