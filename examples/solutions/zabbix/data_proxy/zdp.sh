#!/bin/sh
LOG="/tmp/zdp.log"
NRAPI=""

while true; do
  /usr/local/bin/tsak -name "zdp" -info -log $LOG -nrapi $NRAPI -conf /usr/local/etc/zdp.conf -script /usr/local/etc/zdp.tsak
done
