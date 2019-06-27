# Mandelbrot Set Visualizer - Distributed Systems 
# Install
* You can download the executable for linux
and run it using
```
$ ./executable
```
* Or you can follow these steps:
1) Make sure you have **GoLang** installed on your local machine
   https://golang.org/doc/install
2) After you successfully install **go**, run the following command:
```
$ go get github.com/HakanSunay/mandelbrot
$ mandelbrot
```

## Parameters
  `-c float`
    	Fractal complexity (default 8)
    	
  `-i int`
    	Mandelbrot loop maximum iterations (default 50)
  
  `-o string`
    	Name of the result file (default "zad18.png")
  
  `-r string`
    	Real and Imaginary Number Range (default "-2.0:2.0:-1.0:1.0")
  
  `-s string`
    	Dimensions: width x height (default "640x480")
  
  `-t int`
    	Amount of threads (default 1)

# Results
![](https://github.com/HakanSunay/mandelbrot/blob/master/results/images/2048.png)
