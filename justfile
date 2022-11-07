#!/usr/bin/env just --justfile

# default
default:
  @just --choose

# start wezterm
term:
    wezterm start --class float --cwd={{ invocation_directory() }}

# run go programe
run:
    go run .

update:
  go get -u
  go mod tidy -v

gif:
    go build .
    vhs < ./utils/show.tape
    rm -rf md
    rm -rf qwe.md