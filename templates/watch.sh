#!/bin/bash

# Get current directory
pushd `dirname $0` > /dev/null
SCRIPTPATH=`pwd`
popd > /dev/null

while inotifywait -r -e modify -e close_write $SCRIPTPATH
do
    ./render.sh
done