#!/usr/bin/env bash

OUTPUT=/tmp/wall.png
while [ true ]
do
	pscircle --output=$OUTPUT
	 feh --bg-fill $OUTPUT
	
	sleep 5	
done
