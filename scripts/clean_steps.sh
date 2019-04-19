#!/usr/bin/env bash

for line in $(find . -iname 'main.go'); do

OUT=$(dirname $line)
echo "Removing step binary $OUT/main"
rm $OUT/main
done