#!/usr/bin/env bash
TMP=/tmp/rofi-files-selected
SELECTED_PATH=$(cat $TMP)

if [ "$@" ]
then
	if [ -z "$SELECTED_PATH" ]; then
		echo "$@" > $TMP
		echo "open"
		echo "vim"
	else
		case "$@" in
		"open")
			CMD="xdg-open ${SELECTED_PATH}"
			;;
		"vim")
			CMD="mate-terminal -e \"vim ${SELECTED_PATH}\""
			;;
		esac
		
		coproc ( $CMD > /dev/null 2>&1 )
		rm $TMP
	fi
else
	locate home
fi
