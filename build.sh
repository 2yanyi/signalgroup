#!/bin/sh
npm run build
cp icon/dingchat.png dist/dingchat-linux-x64
cp install dist
printf "打包中... "
tar -Jcf dingchat-linux-x64.tar.xz dist
chmod a+rwx dingchat-linux-x64.tar.xz
printf "OK\n"
rm -r dist
