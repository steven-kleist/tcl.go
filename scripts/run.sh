#!/usr/bin/env sh

TCL="./build/tcl"
if [ $(uname) = "Windows_NT" ]; then
	TCL="${TCL}.exe"
fi

$TCL ./test/test.tcl
