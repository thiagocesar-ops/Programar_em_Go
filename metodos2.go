/*metodos2.go*/

package main

import (
	"fmt"
	"strings"
)

type Pessoa struct 



{
	Nome string
	Sobrenome string 	
}

//Método= oferecer uma funcionalidade a mais para a função.

func (p Pessoa) NomeCompleto() string {
	return p.Nome + " "+ p.Sobrenome

}

func (p *Pessoa) CapitalizarNome() { //Receptor de valor

	p.Nome = strings.ToUpper(p.Nome)
	fmt.Printf("Dentro de CapitalizarNome: %s \n", p.Nome)
}


func main(){

	alguem := Pessoa{"Jose", "de Alencar"} 

	alguem.CapitalizarNome()
	fmt.Println(alguem.NomeCompleto())

}