#!/bin/bash


validateIfDirectoryExists() {

	if [ ! -d "$1" ]; then
		echo "Directory $1 DOES NOT exists." 
		exit 1
	fi
}

validateIfDirectoryExists "$1"

make build -C $1 OUT=$(pwd)