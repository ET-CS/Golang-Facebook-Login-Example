#!/bin/bash

# compile and minify the js files using minifyjs
function compile {
    minifyjs -m --level=1 -i $1 > ../.min/js/$2
}

# Compile but append instead of creating new files.
function append {
    minifyjs -m --level=1 -i $1 >> ../.min/js/$2
}

compile index.js index.min.js
compile vendor/ie10hack.js ie10hack.min.js
