#!/bin/bash
# Github: @ertugrulturan  /  T13R / L1nux-Dev
if ! [ $(id -u) = 0 ]; then
   echo "Installation can only be done as a root user!"
   exit 1
fi
wget https://raw.githubusercontent.com/ertugrulturan/tssh/main/tssh
chmod +x tssh
mv tssh /usr/bin/
clear
echo "Installation is finished. To learn about usage: https://github.com/ertugrulturan/tssh"
