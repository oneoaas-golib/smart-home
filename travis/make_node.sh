#!/usr/bin/env bash

set -o errexit

NODE="/tmp"
TMP_DIR="/tmp"
GOBUILD_LDFLAGS=""
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
common=${BASEDIR}/common.sh ; source "$common" ; if [ $? -ne 0 ] ; then echo "Error - no settings functions $common" 1>&2 ; exit 1 ; fi
GOPATH="${NODE}/vendor"
TMP_DIR="${TMP_DIR}/node"
EXEC=node

main() {

  export DEBIAN_FRONTEND=noninteractive

  if [[ $# = 0 ]] ; then
    echo 'No arguments provided, installing with'
    echo 'default configuration values.'
  fi

  : ${INSTALL_MODE:=stable}

  case "$1" in
    --test)
    __test
    ;;
    --init)
    __init
    ;;
    --clean)
    __clean
    ;;
    --help)
    __help
    ;;
    --build)
    __build
    ;;
    *)
    echo "Error: Invalid argument '$1'" >&2
    __help
    exit 1
    ;;
  esac

}

__test() {

    DIRS=(
        "${NODE}/src/controllers"
        "${NODE}/src/models"
        "${NODE}/src/router"
    )

    for dir in ${DIRS};
    do
        pushd ${BASEDIR}${dir}
        go test -v
        popd
    done
}

__init() {

    mkdir -p ${TMP_DIR}
    cd ${NODE}
    gvt rebuild   
}

__clean() {

    rm -rf ${NODE}/vendor/bin
    rm -rf ${NODE}/vendor/pkg
    rm -rf ${NODE}/vendor/src
    rm -rf ${TMP_DIR}
}

__build() {

    cd ${NODE}/src
    env GOPATH=${GOPATH} go build -ldflags "${GOBUILD_LDFLAGS}" -o ${TMP_DIR}/${EXEC}
    cp ${NODE}/src/node.sample.conf ${TMP_DIR}
}

__help() {
  cat <<EOF
Usage: make_node.sh [options]

Bootstrap Debian 8.0 host with mysql installation.

OPTIONS:

  --test - testing package
  --init - initialize the development environment
  --clean - cleaning of temporary directories
  --build - build backend

  -h / --help - show this help text and exit 0

EOF
}

main "$@"