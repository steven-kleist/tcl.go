#!/usr/bin/env sh


TARGETDIR="./build"
TARGETNAME="tcl"
if [ $(uname) = "Windows_NT" ]; then
	TARGETNAME="${TARGETNAME}.exe"
fi

go build -o "$TARGETDIR/$TARGETNAME" ./cmd/tcl
