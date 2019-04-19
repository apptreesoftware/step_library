#!/usr/bin/env bash

for line in $(find . -iname 'main.go'); do

OUT=$(dirname $line)


#echo "Building step at $line"

go build -o $OUT/main $OUT
done