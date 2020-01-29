#!/bin/bash
#
# Same as `go get` but you can specify a git commit ref.
#
set -e


read -p "This git resets, cleans, builds and tests both unit tests and examples.\n Are you sure? " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    git reset --hard
    git clean -xffd

    make setup && make out
fi