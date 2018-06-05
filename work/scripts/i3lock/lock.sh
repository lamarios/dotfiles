#!/bin/bash
BG=/tmp/screenshot.png
scrot $BG
convert $BG -scale 5% -scale 2000%    $BG
~/scripts/i3lock-color/i3lock \
 --radius=200 \
--ring-width=1 \
--veriftext="" \
--wrongtext="" \
 --indpos="x+(w/2):y+(h/2)" \
-t \
-k \
--timestr="%H:%M" \
--timesize=50 \
--timefont="System San Francisco Display" \
--datefont="System San Francisco Display" \
--datesize=20 \
--timepos="x+(w/2)-(cw/2):y+(h/2)-(ch/2)" \
--timecolor=FFFFFFFF \
--datecolor=FFFFFFAA \
 -i $BG \
--insidevercolor=FFFFFF00 \
--insidewrongcolor=FF000000 \
--insidecolor=FFFFFF00 \
--ringvercolor=FFFFFFFF \
--ringwrongcolor=FF0000AA \
--ringcolor=FFFFFF00 \
--linecolor=FFFFFF00 \
--separatorcolor=FFFFFF00 \
--textcolor=FFFFFFFF \
--keyhlcolor=FFFFFFFF \
--bshlcolor=FFFFFFFF 
