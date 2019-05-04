#!/bin/bash

x=1
while [ $x -le 32 ]
do
  echo `./mandelbrot -r -2.0:2.0:-2.0:2.0 -s 2048x2048 -t $x`
  x=$(( $x + 1 ))
done
