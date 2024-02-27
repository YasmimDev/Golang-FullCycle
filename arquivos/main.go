package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//====== Escreve dados dentro do arquivo criado ======//
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))

	//====== Informar o tamanho do arquivo ======//
	// tamanho, err := f.WriteString("hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes, ", tamanho)

	f.Close()

	//===== Leitura de arquivo ====/

	arq, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arq))


}
