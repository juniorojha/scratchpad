package main

import (
	"fmt"
)

func Sqrt(x float64) (float64) {
z := x/2.0
i:=0
for {
var t, delz float64
t = z
//fmt.Println(t)
z = z - (((z*z)-x)/(2*z))
delz = t - z
//fmt.Println(z, delz)
i+=1

if (delz < 0.000001) {fmt.Println(i)
return z}
}

}

func main() {
	fmt.Println(Sqrt(10000))
}

