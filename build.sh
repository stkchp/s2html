#!/bin/bash
#
# build.sh <input .S file>
#   print html to stdout
#

if [[ $# -ne 1 ]]
then
	echo "Usage: $0 <input .S file>" >&2
	exit 1
fi

if [[ ! -f "$1" ]]
then
	echo "Error: '$1' is not file." >&2
	exit 1
fi

cd $( dirname $0 )
CURDIR=$( pwd )

# build go file
go build s2html.go
[[ $? -ne 0 ]] && echo "Compile Error." >&2 && exit 1

# echo html file
cat $CURDIR/head.html
cat "$1" | $CURDIR/s2html
cat $CURDIR/foot.html

