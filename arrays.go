/* arrays.go*/

package main

import (
	"fmt"
)
func main(){

//Arrays
// 1, 1, 2, 3, 5, 8, 13

	var numeros [7] int // sem numeros ira aparecer 07 zeros
	
// uma das forms de inserir dados em uma array.
	numeros[0] = 1		
	numeros[1] = 1
	numeros[2] = 2
	numeros[3] = 3
	numeros[4] = 5
	numeros[5] = 8
	numeros[6] = 13
	fmt.Println(numeros)  

}