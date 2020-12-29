#!/bin/sh
while true; do
  STATUS=$(( ( RANDOM % 2 )   ))
  PORT=$(( ( RANDOM % 15 )  + 1 ))
  echo "Sending 1.3.6.1.2.1.2.2.1.8.$PORT $STATUS"
  snmptrap -v 1 -c public 127.0.0.1:9162 1.2.3.4 1.2.3.4 3 0 '' 1.3.6.1.2.1.2.2.1.8.$PORT i $STATUS 2> /dev/null
  sleep 1
done
