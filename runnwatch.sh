#!/bin/bash

# first run in background and then watch assets change and update.
# Watch only assets. not the golang server.

./run.sh &
cd templates
./watch.sh