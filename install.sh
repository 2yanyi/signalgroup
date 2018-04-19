#!/bin/sh
APPNAME="dingchat"
echo """[Desktop Entry]
Name = ${APPNAME}
Exec=/usr/local/${APPNAME}
Icon=/usr/local/${APPNAME}.png
Terminal=false
Type=Application
""" > /usr/share/applications/${item}.desktop
