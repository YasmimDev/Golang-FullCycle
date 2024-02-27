package main

import "fmt"


type usuario struct{
	nome string
	idade int
}

func (u usuario) salvar(){
	fmt.Println("salvar no metodo")
}

func main(){
	usuario := usuario{"usuario 1,", 20}
	fmt.Println(usuario)
	usuario.salvar()
}