#!/bin/bash

x=1
while [ $x -le 32 ]
do
  echo `./mandelbrot -t $x`
  x=$(( $x + 1 ))
done
