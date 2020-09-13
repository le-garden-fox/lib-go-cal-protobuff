#!/usr/bin/env bash

function main() {
    case $1 in
    "gen")
        protoc calc/calcpb/calc.proto --go_out=plugins=grpc:.
        ;;
    *)
        echo "Invalid arguments"
        ;;
    esac
}

main "$@"
