# Basic Timer Sercure Shell
Timer SSH - Simple ssh connection script for automatic ping after reboot or during installation
# One Command Install (run root)
```
apt-get install wget curl -y && curl https://raw.githubusercontent.com/ertugrulturan/tssh/main/install.sh | bash
```
# Usage
default connect user = root
default connect port = 22

command:
```
tssh 1.1.1.1
```
special port and user command:
```
tssh 1.1.1.1 user 2222
```
