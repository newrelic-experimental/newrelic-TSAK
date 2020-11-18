#!/bin/sh
snmptrap -e 0x090807060504030201 -v 3 -u "user" -a "MD5" -x "AES" -X "hellohello" -A "hellohello" -l authPriv  127.0.0.1:9163 '' linkUp.0 1.3.6.4.5.6 s "switch is off"
