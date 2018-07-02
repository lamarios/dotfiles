#!/usr/bin/env bash

OUTPUT=/tmp/wall.png
while [ true ]
do
	pscircle --output=$OUTPUT
	wal -i $OUTPUT -o wal-set  && feh --bg-fill $(cat     ~/.cache/wal/wal)
	
	sleep 60
done
