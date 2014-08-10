#!/bin/bash
# build assets, format source code and build

echo Compiling assets...
cd templates && ./render.sh 
cd ..

echo Formatting code...
go fmt main.go

echo Building...
go build main.go
