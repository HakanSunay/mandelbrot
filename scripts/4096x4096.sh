#!/bin/bash

x=1
while [ $x -le 32 ]
do
  echo `./newest -r -2.0:2.0:-2.0:2.0 -s 4096x4096 -t $x`
  x=$(( $x + 1 ))
done
