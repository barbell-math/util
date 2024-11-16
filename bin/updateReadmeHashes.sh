#!/bin/bash

newHash=`git rev-parse origin/main`

findStr="blob/[a-zA-Z0-9]*/"
replaceStr="blob/$newHash/"

grep -rl --exclude-dir=".git" "$findStr" . | xargs sed -i "s#$findStr#$replaceStr#g"
