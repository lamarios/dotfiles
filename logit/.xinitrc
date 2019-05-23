source ~/.profile


# disabling page up and down as the location of those keys is really bad

eval $(/usr/bin/gnome-keyring-daemon --start --components=gpg,pkcs11,secrets,ssh)
export $(gnome-keyring-daemon --start --components=pkcs11,secrets,ssh,gpg)
dbus-update-activation-environment --systemd DISPLAY

xrandr --dpi 96
exec i3
#exec mate-session


