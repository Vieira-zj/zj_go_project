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
    cat .gitmodules
}

function add_sub_git_repos2() {
    git submodule add https://github.com/sergi/go-diff.git src/github.com/sergi/go-diff
    git submodule add https://github.com/BurntSushi/toml.git src/github.com/BurntSushi/toml
    git submodule add https://github.com/golang/lint.git src/golang.org/x/lint
    git submodule add https://github.com/golang/mod.git src/golang.org/x/mod
    git submodule add https://github.com/golang/sync.git src/golang.org/x/sync
    git submodule add https://github.com/golang/tools.git src/golang.org/x/tools
    git submodule add https://github.com/golang/xerrors.git src/golang.org/x/xerrors
    git submodule add https://github.com/dominikh/go-tools.git src/honnef.co/go/tools
    git submodule add https://github.com/mvdan/xurls.git src/mvdan.cc/xurls
}

function install_vscode_plugins() {
    ext_tools=("github.com/mdempsky/gocode" "github.com/ramya-rao-a/go-outline" \
        "github.com/acroca/go-symbols" "golang.org/x/tools/cmd/guru" \
        "golang.org/x/tools/cmd/gorename" "github.com/stamblerre/gocode" \
        "github.com/sqs/goreturns" "golang.org/x/lint/golint"
    )
    
    for tool in ${ext_tools[*]}; do
        echo "install tool: ${tool}"
        cd ${HOME}/Workspaces/zj_go_project/src/$tool
        go install
    done
}

# main
# install_ext_tools
# fix_sub_repo

set +ex
